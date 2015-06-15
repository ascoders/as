'use strict';

define("gamePage", ['jquery', 'marked', 'prettify', 'editor', 'jquery.qrcode', 'jquery.cookie', 'jquery.timeago', 'jquery.autocomplete'], function ($, marked, prettify) {
	var vm = avalon.define({
		$id: "gamePage",
		type: 0, //父级分类类型
		topic: {}, // 文章信息
		replys: [], // 回复 嵌套评论整合在其中
		same: [], //类似文章
		pagin: '', //分页
		inputContent: '', //编辑框内容
		inputEdit: false, //是否为编辑状态
		inputReply: -1, //当前输入框评论的回复
		addTagInput: '', //新增标签输入框内容
		tag: false, //是否在编辑标签状态
		rTag: false, //是否在删除标签状态
		isManagers: false, //是否为管理员组
		temp: { //存储数据
			editor: {}, //编辑器
			from: 0, //显示起始位置
			number: 0, //存储每页显示多少回复
			freshSame: function () { // 获取类似类容
				post('/api/tag/same', {
					game: avalon.vmodels.gameBase.game.Id,
					id: vm.topic.Id
				}, null, '', function (data) {
					data = data || [];

					for (var key = 0; key < data.length; key++) {
						//如果没有标题，取内容
						if (data[key].Title == "") {
							data[key].Title = subStr(data[key].Content, 0, 25);
						}

						// 取出内容中&nbsp;
						data[key].Content = data[key].Content.replace(/&nbsp;/g, '');
					}

					vm.same = data || [];
				});
			}
		},
		changeCategory: function () { // 修改所属分类
			var category = $(this).val();
			post('/api/game/changeCategory', {
				topic: vm.topic.Id,
				category: category
			}, '移动成功', '分类移动失败：');
		},
		rendered: function () { // 回复加载完毕
			jbox();
		},
		operate: function (type, id) { //操作
			switch (type) {
			case 'top': //置顶
				var _type = type;
				if (vm.topic.Top > 0) {
					_type = 'topCancle';
				}
				post('/api/game/operate', {
					id: id,
					type: _type
				}, '操作成功', '', function (data) {
					vm.topic.Top = data;
				});
				break;
			case 'good': //加精
				var _type = type;
				if (vm.topic.Good) {
					_type = 'goodCancle';
				}
				post('/api/game/operate', {
					id: id,
					type: _type
				}, '操作成功', '', function (data) {
					vm.topic.Good = !vm.topic.Good;
				});
				break;
			case 'delete': //删除
				confirm('删除此帖吗？', function () {
					post('/api/game/operate', {
						id: id,
						type: type
					}, '已删除', '', function (data) {
						// 根据分类路径名获取id
						var categoryPath = '';
						for (var key in avalon.vmodels.gameBase.categorys.$model) {
							if (avalon.vmodels.gameBase.categorys.$model[key].Id == vm.topic.Category) {
								categoryPath = avalon.vmodels.gameBase.categorys.$model[key].Category;
								break;
							}
						}

						//跳转回该分类首页
						avalon.router.navigate('/g/' + avalon.vmodels.gameBase.game.Id + '/' + categoryPath);
					});
				});
				break;
			}
		},
		deleteReply: function (index) { // 删除回复
			confirm('删除此回复吗？', function () {
				post('/api/game/operate', {
					reply: vm.replys[index].Id,
					type: 'deleteReply'
				}, '已删除', '', function (data) {
					vm.replys.removeAt(index);
				});
			});
		},
		submit: function () { // 回复 编辑 嵌套评论
			//判断当前输入框状态
			var type = 'reply'; //回复
			if (vm.inputEdit) { //编辑状态
				type = 'edit';
			}
			if (vm.inputReply != -1) { //嵌套回复状态
				type = 'rReply';
			}

			var replyId = '';
			if (vm.inputReply != -1) {
				replyId = vm.replys[vm.inputReply].Id;
			}

			if (vm.inputContent.length < 3) {
				notice('内容至少3个字符', 'red');
				return;
			}

			post('/api/game/addTopic', {
				game: avalon.vmodels.gameBase.game.Id,
				topic: vm.topic.Id, // 文章id
				category: vm.topic.Category, // 所属分类 后台需要验证权限
				content: vm.inputContent,
				type: type,
				reply: replyId // 嵌套评论参数，表示评论给哪个回复
			}, null, '', function (data) {

				//清空内容
				vm.inputTitle = "";
				vm.inputContent = "";
				vm.temp.editor.freshPreview();

				// 跳转前缀
				var jumpPrefix = '/g/' + avalon.vmodels.gameBase.game.Id + '/' + vm.topic.Id;

				if (vm.type == 1) {
					jumpPrefix = '/g/' + avalon.vmodels.gameBase.game.Id + '/' + avalon.vmodels.gameBase.category + '/' + vm.topic.Id;
				}

				switch (type) {
				case 'reply':
					//跳转到最后一页
					avalon.router.navigate(jumpPrefix +
						'?number=' + vm.temp.number + '&from=' +
						Math.floor(parseInt(vm.topic.OutReply) / vm.temp.number) * vm.temp.number, {
							reload: true
						});
					break;
				case 'edit':
					//刷新本页
					avalon.router.navigate(jumpPrefix +
						'?number=' + vm.temp.number + '&from=' +
						vm.temp.from, {
							reload: true
						});

					//滑到顶部
					$("html, body").animate({
						scrollTop: 0
					}, 300);
					break;
				case 'rReply':
					break;
				}
			});
		},
		toggleEdit: function () { //切换/取消编辑状态
			if (vm.inputReply != -1) {
				//先取消嵌套评论模式
				vm.toggleReply(vm.inputReply);
			}

			vm.inputReply = -1;
			vm.inputEdit = !vm.inputEdit;
			if (vm.inputEdit) {
				vm.inputContent = vm.topic.Content;

				//同步review
				vm.temp.editor.freshPreview();

				// 在文档模式下，改变滑动位置
				if (vm.type == 1) {
					//滑到编辑区
					$('.paper').animate({
						scrollTop: $("#game-page #text").offset().top - 122 - 32
					}, 300);
				} else {
					//滑到编辑区
					$('html ,body').animate({
						scrollTop: $("#game-page #text").offset().top - 32 - 0.1
					}, 300);
				}


			} else {
				vm.inputContent = "";
			}
			//刷新预览区
			vm.temp.editor.freshPreview();
		},
		toggleReply: function (index) { // 切换/取消嵌套回复状态
			if (vm.inputEdit) {
				//先取消编辑模式
				vm.toggleEdit();
			}

			//确定移动到的index,给inputReply赋值
			if (vm.inputReply == index) {
				vm.inputReply = -1;
			} else {
				vm.inputReply = index;
			}

			//进入嵌套回复 / 取消
			if (vm.inputReply != -1) { // 进入嵌套回复
				// 编辑器取消markdown功能
				vm.temp.editor.enableMarkdown(false);

				//把输入框追加到此回复之后
				$("#game-page #editContent").removeClass('g-bd editContentAlone').addClass('editContentInline');
				$('#game-page .reply:eq(' + index + ')').after($("#game-page #editContent"));
			} else { // 取消
				// 编辑器恢复markdown功能
				vm.temp.editor.enableMarkdown(true);

				//输入框恢复到主题框之后
				var rmclass = 'g-bd editContentAlone';
				if (vm.type == 1) {
					rmclass = 'editContentAlone';
				}
				$("#game-page #editContent").addClass(rmclass).removeClass('editContentInline');
				$('#game-page #content').after($("#game-page #editContent"));
			}
			// 刷新预览区
			vm.temp.editor.freshPreview();
		},
		toggleTag: function () { // 新增标签输入框组是否显示
			vm.tag = !vm.tag;

			// 获取焦点
			$('#game-page #tag-input').focus();
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

			if (vm.topic.Tag.size() >= 5) {
				notice('最多5个标签', 'red');
				return;
			}

			//是否重复
			if ($.inArray(vm.addTagInput, vm.topic.Tag.$model) != -1) {
				notice('标签不能重复', 'red');
				return;
			}

			post('/api/tag/bind', {
				topic: vm.topic.Id,
				name: vm.addTagInput
			}, '标签已添加', '', function (data) {
				vm.topic.Tag.push(vm.addTagInput);

				// 输入框置空
				vm.addTagInput = '';

				// 取消新增状态
				vm.tag = false;

				// 获取类似文章
				vm.temp.freshSame();
			});
		},
		jumpOrRemove: function (name) {
			if (vm.rTag) {
				post('/api/tag/unBind', {
					topic: vm.topic.Id,
					name: name
				}, '标签已删除', '', function (data) {
					vm.topic.Tag.remove(name);

					// 获取类似文章
					vm.temp.freshSame();
				});
			} else {
				avalon.router.navigate('/g/' + avalon.vmodels.gameBase.game.Id + '/tag?tag=' + name);
			}
		},
		toggleRemoveTag: function () { //切换删除标签状态
			vm.rTag = !vm.rTag;
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			$.when(avalon.vmodels.global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
				if ($.inArray(avalon.vmodels.global.my.id, avalon.vmodels.gameBase.game.Managers) > -1) {
					vm.isManagers = true;
				}
			});

			if (avalon.vmodels.gameListDoc != undefined) {
				avalon.vmodels.gameListDoc.$pageReady(param);
			}

			//如果是编辑状态，取消
			if (vm.inputEdit) {
				vm.toggleEdit();
			}

			//如果是嵌套回复状态，取消
			if (vm.inputReply != -1) {
				vm.toggleReply(vm.inputReply);
			}

			vm.tag = false;
			vm.rTag = false;

			mmState.query.from = mmState.query.from || 0;
			mmState.query.number = mmState.query.number || 20;

			vm.temp.from = mmState.query.from;
			vm.temp.number = mmState.query.number;

			//获取内容
			post('/api/game/getPage', {
				game: avalon.vmodels.gameBase.game.Id,
				id: param.id,
				from: mmState.query.from,
				number: mmState.query.number
			}, null, '获取信息失败：', function (data) {
				// 改变页面titile
				document.title = data.topic.Title + ' - ' + avalon.vmodels.gameBase.game.Name;

				//用户头像路径处理
				data.topic.AuthorImage = userImage(data.topic.AuthorImage);

				//内容markdown处理
				data.topic.MarkedContent = data.topic.Content == '' ? '暂无内容' : marked(data.topic.Content);
				vm.topic = data.topic || {};

				avalon.nextTick(function () {
					//代码高亮
					$('pre').addClass('prettyprint pre-scrollable linenums');
					prettify.prettyPrint();

					//timeago
					$('.timeago').timeago();
				});

				//处理回复内容
				for (var key in data.replys) {
					if (!data.replys.hasOwnProperty(key) || key === "hasOwnProperty") {
						continue;
					}

					data.replys[key].AuthorImage = userImage(data.replys[key].AuthorImage);
					data.replys[key].Content = marked(data.replys[key].Content);
					data.replys[key].ReplyCache = data.replys[key].ReplyCache || [];
					for (var k in data.replys[key].ReplyCache) {
						data.replys[key].ReplyCache[k].AuthorImage = userImage(data.replys[key].ReplyCache[k].AuthorImage);
					}
				}
				vm.replys = data.replys || [];

				// 生成分页
				vm.pagin = createPagin(mmState.query.from, mmState.query.number, data.count);

				// 获取类似文章
				vm.temp.freshSame();
			});
		}

		$ctrl.$onRendered = function () {
			// 实例化markdown编辑器
			vm.temp.editor = new $("#game-page #text").MarkEditor({
				uploadUrl: "/api/game/upload",
				uploadParams: {
					type: 'gameUserImage'
				}
			});

			//刷新编辑器dom
			vm.temp.editor.createDom();

			//jbox显示提示
			jbox();

			//显示二维码
			$("#j-qrcode").qrcode({
				render: "canvas", //canvas or table方式
				width: 150, //宽度
				height: 150, //高度
				background: '#fff',
				foreground: '#333',
				text: window.location.href //当前url
			});

			//图片最宽100%
			$('#game-page img').css('max-width', '100%');

			//获取xsrftoken
			var xsrf = $.cookie("_xsrf");
			if (!xsrf) {
				return;
			}
			var xsrflist = xsrf.split("|");
			var xsrftoken = Base64.decode(xsrflist[0]);

			// 搜索标签自动完成
			$('#game-page #tag-input').autocomplete({
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

			/*
						var dropbox = document.getElementById('text');

						dropbox.addEventListener('drop', drop, false);

						function drop(evt) {
							evt.stopPropagation();
							evt.preventDefault();
							var imageUrl = evt.dataTransfer.getData('URL');
							alert(imageUrl);
						}
						*/
		}
	});
});