'use strict';

define("gameAdminCategory", ['jquery', 'jquery.jbox'], function ($) {

	var vm = avalon.define({
		$id: "gameAdminCategory",
		info: '', //当前在菜单的位置
		add: { //新增分类参数
			name: '', //名称
			path: '', //路径
			number: 5, //推荐数量
			add: 1, //发帖限制
			reply: 1, //回帖限制
			type: 0 //类型
		},
		$temp: {
			changeRecommendPri: function (category, value) { //改变优先级
				return post('/api/game/changeRecommendPri', {
					game: avalon.vmodels.gameBase.game.Id,
					category: category,
					value: value
				}, null, '');
			},
			addModal: null
		},
		changeLay: function (index, type) { //改变层级 由上到下优先级逐渐减小
			// 新建两个promise
			var defer1, defer2 = $.Deferred();

			if (type == 1) { // 上移
				if (index == 0) {
					return;
				}

				// 如果层级与上个相同,此优先级-1
				if (avalon.vmodels.gameBase.categorys[index].RecommendPri == avalon.vmodels.gameBase.categorys[index - 1].RecommendPri) {
					defer1 = vm.$temp.changeRecommendPri(avalon.vmodels.gameBase.categorys[index].Category, avalon.vmodels.gameBase.categorys[index].RecommendPri - 1);
					defer2.resolve();
					avalon.vmodels.gameBase.categorys[index].RecommendPri++;
				} else { // 层级不同,则对调层级
					defer1 = vm.$temp.changeRecommendPri(avalon.vmodels.gameBase.categorys[index].Category, avalon.vmodels.gameBase.categorys[index - 1].RecommendPri);
					defer2 = vm.$temp.changeRecommendPri(avalon.vmodels.gameBase.categorys[index - 1].Category, avalon.vmodels.gameBase.categorys[index].RecommendPri);
					var $temp = avalon.vmodels.gameBase.categorys[index].RecommendPri;
					avalon.vmodels.gameBase.categorys[index].RecommendPri = avalon.vmodels.gameBase.categorys[index - 1].RecommendPri;
					avalon.vmodels.gameBase.categorys[index - 1].RecommendPri = $temp;
				}
			} else { // 下移
				if (index == avalon.vmodels.gameBase.categorys.size() - 1) {
					return;
				}

				// 如果层级与下个相同,此优先级+1
				if (avalon.vmodels.gameBase.categorys[index].RecommendPri == avalon.vmodels.gameBase.categorys[index + 1].RecommendPri) {
					defer1 = vm.$temp.changeRecommendPri(avalon.vmodels.gameBase.categorys[index].Category, avalon.vmodels.gameBase.categorys[index].RecommendPri + 1);
					defer2.resolve();
					avalon.vmodels.gameBase.categorys[index].RecommendPri--;
				} else { // 层级不同,则对调层级
					defer1 = vm.$temp.changeRecommendPri(avalon.vmodels.gameBase.categorys[index].Category, avalon.vmodels.gameBase.categorys[index + 1].RecommendPri);
					defer2 = vm.$temp.changeRecommendPri(avalon.vmodels.gameBase.categorys[index + 1].Category, avalon.vmodels.gameBase.categorys[index].RecommendPri);
					var $temp = avalon.vmodels.gameBase.categorys[index].RecommendPri;
					avalon.vmodels.gameBase.categorys[index].RecommendPri = avalon.vmodels.gameBase.categorys[index + 1].RecommendPri;
					avalon.vmodels.gameBase.categorys[index + 1].RecommendPri = $temp;
				}
			}

			$.when(defer1, defer2).done(function () {
				// 重排序
				avalon.vmodels.gameBase.categorys.sort(function (a, b) {
					return a.RecommendPri > b.RecommendPri ? 1 : -1;
				});
			});
		},
		openAddModal: function () { // 打开添加模态框
			vm.$temp.addModal.open();
		},
		hasChange: function (index) { // 用户修改后，保存按钮变亮
			$('#game-admin-category #save' + index).removeClass('btn-default').addClass('btn-success');
		},
		save: function (index) { // 保存分类更新
			$(this).removeClass('btn-success').addClass('btn-default');

			post('/api/game/updateCategory', {
				game: avalon.vmodels.gameBase.game.Id,
				id: avalon.vmodels.gameBase.categorys[index].Id,
				path: avalon.vmodels.gameBase.categorys[index].Category,
				name: avalon.vmodels.gameBase.categorys[index].CategoryName,
				number: avalon.vmodels.gameBase.categorys[index].Recommend,
				add: avalon.vmodels.gameBase.categorys[index].Add,
				reply: avalon.vmodels.gameBase.categorys[index].Reply,
				_type: avalon.vmodels.gameBase.categorys[index].Type
			}, '保存成功', '');
		},
		deleteCategory: function (index) { // 删除分类
			confirm('删除分类 ' + avalon.vmodels.gameBase.categorys[index].CategoryName + ' 吗', function () {
				post('/api/game/deleteCategory', {
					game: avalon.vmodels.gameBase.game.Id,
					id: avalon.vmodels.gameBase.categorys[index].Id
				}, '已删除', '', function () {
					// 强制刷新本页
					avalon.router.navigate(window.location.pathname, {
						reload: true
					});
				});
			});
		}
	});

	return avalon.controller(function ($ctrl) {

		$ctrl.$onEnter = function (param, rs, rj) {
			// 改变页面titile
			document.title = '分类管理 - 管理 - ' + avalon.vmodels.gameBase.game.Name;

			// 设置状态
			avalon.vmodels.gameAdmin.info = 'category';
		}

		$ctrl.$onRendered = function () {
			// jbox提示
			jbox();

			// 初始化弹窗
			if ($.isEmptyObject(vm.$temp.addModal)) {
				vm.$temp.addModal = new jBox('Confirm', {
					minWidth: '200px',
					content: $('#game-admin-category #modal'),
					animation: 'flip',
					confirmButton: '确定',
					cancelButton: '取消',
					overlay: false,
					confirm: function () {
						//发送请求
						post('/api/game/addCategory', {
							game: avalon.vmodels.gameBase.game.Id,
							name: vm.add.name,
							path: vm.add.path,
							number: vm.add.number,
							add: vm.add.add,
							reply: vm.add.reply,
							_type: vm.add.type,
							pri: avalon.vmodels.gameBase.categorys.size()
						}, '新增成功', '新增失败：', function () { //新增成功
							// 强制刷新本页
							avalon.router.navigate(window.location.pathname, {
								reload: true
							})
						})
					}
				})
			}
		}

	});

});