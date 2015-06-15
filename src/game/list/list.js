'use strict';

define("gameList", ['jquery', 'editor', 'jquery.autocomplete'], function ($) {
	var vm = avalon.define({
		$id: "gameList",
		lists: [], //列表信息
		pagin: '', //分页
		inputTitle: '', //准备提交的标题
		inputContent: '', //准备提交的内容
		loading: false, //是否在loading状态
		isManagers: false, //是否为管理员组
		addTagInput: '', //新增标签输入框内容
		tag: false, //是否在编辑标签状态
		searchTag: '', //当前搜索的标签（仅在category:tag路由下有效）
		hotTags: [], //热门标签（仅在category:tag路由下有效）
		tagArray: [], //准备发布文章的标签数组
		$temp: { //存储数据
			editor: null, //编辑器
			from: 0, //显示起始位置
			number: 0 //存储每页显示多少回复
		},
		toggleTag: function () { // 新增标签输入框组是否显示
			vm.tag = !vm.tag;

			// 获取焦点
			$('#game-list #tag-input').focus();
		},
		addTag: function () { // 新增标签
			if (vm.addTagInput == '') {
				notice('标签不能为空', 'red');
				return;
			}

			if (vm.addTagInput.length > 15) {
				notice('标签最大长度为15', 'red');
				return;
			}

			if (vm.tagArray.size() >= 5) {
				notice('最多5个标签', 'red');
				return;
			}

			//是否重复
			if ($.inArray(vm.addTagInput, vm.tagArray.$model) != -1) {
				notice('标签不能重复', 'red');
				return;
			}

			vm.tagArray.push(vm.addTagInput);

			// 输入框置空
			vm.addTagInput = '';

			// 取消新增状态
			vm.tag = false;
		},
		RemoveTag: function (name) {
			vm.tagArray.remove(name);
		},
		submit: function () { //提交发帖	
			if (vm.inputContent.length < 3) {
				notice('内容至少3个字符', 'red');
				return;
			}

			// 按钮变灰
			$(this).attr('disabled', "disabled").text('提交中..');

			// 根据分类路径名获取id
			var categoryId = '';
			for (var key in avalon.vmodels.gameBase.categorys.$model) {
				if (avalon.vmodels.gameBase.categorys.$model[key].Category == avalon.vmodels.gameBase.category) {
					categoryId = avalon.vmodels.gameBase.categorys.$model[key].Id;
					break;
				}
			}

			post('/api/game/addTopic', {
				game: avalon.vmodels.gameBase.game.Id,
				category: categoryId,
				title: vm.inputTitle,
				content: vm.inputContent,
				tag: vm.tagArray.$model,
				type: 'addTopic' //发帖
			}, '发布成功', '', function (data) {
				//跳转到帖子页面
				avalon.router.navigate('/g/' + avalon.vmodels.gameBase.game.Id + '/' + data);

				//清空内容
				vm.inputTitle = "";
				vm.inputContent = "";
				vm.$temp.editor.freshPreview();
			});
		}
	});
	return avalon.controller(function ($ctrl) {

		$ctrl.$vmodels = [vm];

		$ctrl.$onEnter = function (param, rs, rj) {
			$.when(avalon.vmodels.global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
				if ($.inArray(avalon.vmodels.global.my.id, avalon.vmodels.gameBase.game.Managers) > -1) {
					vm.isManagers = true;
				}
			});

			// 初始化
			vm.tag = false;
			vm.searchTag = '';

			mmState.query.from = mmState.query.from || 0;
			mmState.query.number = mmState.query.number || 20;

			vm.$temp.from = mmState.query.from;
			vm.$temp.number = mmState.query.number;

			// 赋值category
			avalon.vmodels.gameBase.category = param.category;

			// 赋值base当前所在位置
			avalon.vmodels.gameBase.baseRouter = param.category;

			// 如果列表有内容，则设置loading长宽
			if (vm.lists.size() > 0 && mmState.prevState.stateName == mmState.currentState.stateName) {
				$('#game-list #loading').css({
					width: $('#game-list #main-content').width() + 'px',
					height: $('#game-list #main-content').height() + 'px'
				});

				// 显示加载模块
				vm.loading = true;
			}

			// 根据分类路径名获取id
			var categoryId = '';
			for (var key in avalon.vmodels.gameBase.categorys.$model) {
				if (avalon.vmodels.gameBase.categorys.$model[key].Category == param.category) {
					categoryId = avalon.vmodels.gameBase.categorys.$model[key].Id;

					// 改变页面titile
					document.title = avalon.vmodels.gameBase.categorys.$model[key].CategoryName + ' - ' + avalon.vmodels.gameBase.game.Name;
					break;
				}
			}

			var postUrl = '/api/game/getList';
			var postParams = {
				game: avalon.vmodels.gameBase.game.Id,
				category: categoryId,
				from: mmState.query.from,
				number: mmState.query.number
			};

			// 如果是标签页
			if (mmState.query.tag != undefined) {
				var postUrl = '/api/tag/getList';
				postParams = {
					game: avalon.vmodels.gameBase.game.Id,
					tag: mmState.query.tag,
					from: mmState.query.from,
					number: mmState.query.number
				};

				// 赋值tag标签
				vm.searchTag = mmState.query.tag || "";

				// 请求热门标签
				if (vm.hotTags.size() == 0) {
					post('/api/tag/hot', {
						game: avalon.vmodels.gameBase.game.Id
					}, null, '热门标签获取失败', function (data) {
						vm.hotTags = data;
					});
				}

				// 改变页面titile
				$('title').text('标签：' + mmState.query.name + ' - ' + avalon.vmodels.gameBase.game.Name);
			}

			//请求获取分页信息
			post(postUrl, postParams, null, '获取信息失败：', function (data) {
				data.articles = data.articles || [];
				for (var key = 0; key < data.articles.length; key++) {
					//用户头像处理
					data.articles[key].AuthorImage = userImage(data.articles[key].AuthorImage);
					//图片附着在articles对象上
					data.articles[key].Images = data.images[key];

					// 取出内容中&nbsp;
					data.articles[key].Content = data.articles[key].Content.replace(/&nbsp;/g, '');
				}
				vm.lists = data.articles || [];

				avalon.nextTick(function () {
					// 加载完毕
					vm.loading = false;
				});

				//生成分页
				vm.pagin = createPagin(mmState.query.from, mmState.query.number, data.count);
			});
		}

		$ctrl.$onRendered = function () {
			// 实例化markdown编辑器
			vm.$temp.editor = new $("#game-list #text").MarkEditor({
				uploadUrl: "/api/game/upload",
				uploadParams: {
					game: avalon.vmodels.gameBase.game.Id,
					uploadType: 'userImage'
				}
			});

			//刷新编辑器dom
			vm.$temp.editor.createDom();

			//获取xsrftoken
			var xsrf = $.cookie("_xsrf");
			if (!xsrf) {
				return;
			}
			var xsrflist = xsrf.split("|");
			var xsrftoken = Base64.decode(xsrflist[0]);

			// 搜索标签自动完成
			$('#game-list #tag-input').autocomplete({
				serviceUrl: '/api/tag/searchTag',
				type: 'post',
				deferRequestBy: 300,
				params: {
					_xsrf: xsrftoken,
					game: avalon.vmodels.gameBase.game.Id
				},
				onSelect: function (suggestion) {
					vm.addTagInput = suggestion.value;
					vm.addTag();
				}
			});

			jbox();
		}

	});
});