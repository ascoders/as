'use strict';

define("gameListDoc", ['jquery', 'jquery.contextMenu', 'jquery.ui'], function ($) {
	var vm = avalon.define({
		$id: "gameListDoc",
		isManagers: false, //是否为管理员组
		tree: [], // 左侧菜单树
		$mouseObj: {}, // 鼠标右键时，操作的节点
		$mouseDownObj: {}, // 鼠标左键刚按下时，操作的节点
		nowTopicId: '', // 当前文章id
		$categoryId: '', // 当前分类id
		$treeId: '', // 根目录id
		docMouseUp: function (event, obj) { // 鼠标松开
			if (event.button == 2) { // 右键
				vm.$mouseObj = obj;
			}
		},
		docMouseDown: function (event, obj) { // 鼠标点击
			if (event.button == 0) { // 左键
				vm.$mouseDownObj = obj;
			}
		},
		add: { // 新增
			show: false, // 显示新增
			type: '', // 添加类型 文档 or 文件
			title: '',
			$parentId: '', // 当前父级id
			$docId: '',
			titleBlur: function () { // 保存标题
				if (vm.add.title == '') {
					vm.add.show = false;
					return;
				}

				var isFolder = false;
				if (vm.add.type == 'folder') {
					isFolder = true;
				}

				post('/api/game/addTopic', {
					game: avalon.vmodels.gameBase.game.Id,
					category: vm.$categoryId,
					title: vm.add.title,
					type: 'addTopic', //发帖
					isFolder: isFolder,
					docId: vm.add.$docId,
					docParent: vm.add.$parentId
				}, '新建成功', '', function (data) {
					var parent, parentId;
					// 在根目录添加
					if ($.isEmptyObject(vm.$mouseObj)) {
						parent = null;
					} else {
						// 在子目录添加
						parent = vm.$mouseObj;
					}

					var index = parent == null ?
						vm.tree.size() : parent.child.size();

					var obj = {
						id: data,
						name: vm.add.title,
						child: [],
						_folder: isFolder,
						_name: vm.add.title,
						_show: true,
						_active: false,
						_ajax: false,
						_loading: false,
						_expend: false,
						$parent: parent,
						$index: index
					};

					// 追加
					if ($.isEmptyObject(vm.$mouseObj)) {
						vm.tree.push(obj);
					} else { // 子目录 
						// 如果没有ajax则ajax【不追加】
						if (!vm.$mouseObj._ajax) {
							vm.getChild(vm.$mouseObj);
						} else { // ajax过的，直接展开并追加
							// 展开目录
							vm.$mouseObj._expend = true;

							vm.$mouseObj.child.push(obj);
						}
					}

					// 清空添加
					vm.add.show = false;
					vm.add.title = '';

					vm.$mouseObj = {};
				}, function () {
					vm.add.show = false;
					vm.add.title = '';

					vm.$mouseObj = {};
				});
			},
			keyUp: function (event) { // 回车键新增
				if (event.keyCode == "13") {
					vm.add.titleBlur();
				}
			}
		},
		getChild: function (obj) {
			// 如果在拖拽中，强制关闭
			if ($(this).attr('dragging')) {
				obj._expend = false;
				return;
			}

			obj._expend = !obj._expend;

			if (obj._ajax) { // 已经查询过信息
				return;
			}

			obj._loading = true;

			post('/api/doc/getDoc', {
				id: obj.id
			}, null, null, function (data) {
				obj._ajax = true;

				for (var i = 0, j = data.cs.length; i < j; i++) {
					obj.child.push({
						id: data.cs[i]._id,
						name: data.cs[i].n,
						child: [],
						_folder: data.cs[i].ifr,
						_name: data.cs[i].n,
						_show: true,
						_active: false,
						_ajax: false,
						_expend: false,
						_loading: false,
						$parent: obj,
						$index: i
					});
				}

				avalon.nextTick(function () {
					vm.$sortable();
				});

				obj._loading = false;
			}, function () {
				obj._ajax = true;
				obj._loading = false;
			});
		},
		$getParent: false, // 是否获取过所有父级
		$pageReady: function (param) { // 子文章页面加载完毕
			vm.nowTopicId = param.id;

			// 页面加载后第一次获取父级
			if (vm.$getParent) {
				return;
			}
			vm.$getParent = true;

			post('/api/doc/parents', {
				id: param.id
			}, null, null, function (data) {
				if (data.length == 0) { // 根目录下页面
					for (var j = vm.tree.$model.length - 1; j >= 0; j--) {
						if (vm.tree.$model[j].id == param.id) {
							vm.tree[j]._active = true;
						}
					}
				} else {
					// 倒序循环tree的下标
					var i = data.length - 1;

					var handleTree = function (tree) {
						// 寻找tree中需要添加的子节点
						var targetTree;

						// 添加到指定分支上
						for (var j = 0, k = tree.size(); j < k; j++) {
							if (tree[j].id == data[i]._id) {
								tree[j]._ajax = true;
								tree[j]._expend = true;

								targetTree = tree[j];
							}
						}

						// 为数组赋值
						for (var j = 0, k = data[i].cs.length; j < k; j++) {
							targetTree.child.push({
								id: data[i].cs[j]._id,
								name: data[i].cs[j].n,
								child: [],
								_folder: data[i].cs[j].ifr,
								_name: data[i].cs[j].n,
								_show: true,
								_active: false,
								_ajax: false,
								_loading: false,
								_expend: false,
								$parent: targetTree,
								$index: j
							});
						}

						i--;

						if (i < 0) {
							return;
						}

						handleTree(targetTree.child);
					}

					handleTree(vm.tree);

					avalon.nextTick(function () {
						vm.$sortable();
					});
				}
			});
		},
		$sortable: function () {
			var startIndex;

			$(".draggable").sortable({
				containment: ".menu",
				items: "> li",
				scroll: false,
				distance: 5,
				opacity: 0.7,
				placeholder: "ui-state-highlight",
				cursor: "move",
				zIndex: 1,
				start: function (event, ui) {
					startIndex = ui.item.index();

					$(event.toElement).attr('dragging', 'dragging')
					$(event.toElement).parent().addClass('dragging');
				},
				stop: function (event, ui) {
					$(event.toElement).removeAttr('dragging');
					$(event.toElement).parent().removeClass('dragging');
				},
				update: function (event, ui) { // 位置发生改变
					var parentArray = vm.$mouseDownObj.$parent ?
						vm.$mouseDownObj.$parent.child : vm.tree;

					var parentId = vm.$mouseDownObj.$parent ?
						vm.$mouseDownObj.$parent.id : vm.$categoryId;

					var endIndex = ui.item.index();

					post('/api/doc/exchange', {
						id: parentId,
						from: startIndex,
						to: endIndex
					}, null, '', function () {
						/*
												// 此位置之前的组件，下标依次减少
												if (startIndex < endIndex) {
													for (var i = startIndex + 1; i <= endIndex; i++) {
														parentArray[i].$index = parentArray[i].$index - 1;
													}
												} else {
													for (var i = startIndex - 1; i >= endIndex; i--) {
														parentArray[i].$index = parentArray[i].$index + 1;
													}
												}

												// 更新拖拽组件位置
												parentArray[startIndex].$index = endIndex;

												// 重新排序
												parentArray.sort(function compare(a, b) {
													return a.$index > b.$index;
												});
						*/

						/*
												var temp = $.extend({}, parentArray[startIndex].$model);
												console.log(parentArray);

												// 此位置之前的组件，下标依次减少
												if (startIndex < endIndex) {
													for (var i = startIndex; i < endIndex; i++) {
														parentArray.set(i, parentArray[i + 1].$model);
													}
												} else {
													for (var i = startIndex; i > endIndex; i--) {
														parentArray.set(i, parentArray[i - 1].$model);
													}
												}

												//parentArray.set(endIndex, temp);
												*/



					});
				}
			});

			$(".draggable").disableSelection();
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			$.when(avalon.vmodels.global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
				if ($.inArray(avalon.vmodels.global.my.id, avalon.vmodels.gameBase.game.Managers) > -1) {
					avalon.vmodels.gameList.isManagers = true;
				}
			});

			// 赋值category
			avalon.vmodels.gameBase.category = param.category;

			// 赋值base当前所在位置
			avalon.vmodels.gameBase.baseRouter = param.category;

			// 生成左侧树状列表（根目录）

			// 根据分类路径名获取id
			for (var key in avalon.vmodels.gameBase.categorys.$model) {
				if (avalon.vmodels.gameBase.categorys.$model[key].Category == avalon.vmodels.gameBase.category) {
					vm.$categoryId = avalon.vmodels.gameBase.categorys.$model[key].Id;
					break;
				}
			}

			vm.tree.clear();

			// 获取根目录信息
			post('/api/doc/getDoc', {
				id: vm.$categoryId
			}, null, null, function (data) {
				var tree = [];

				vm.$treeId = data._id;

				for (var i = 0, j = data.cs.length; i < j; i++) {
					tree.push({
						id: data.cs[i]._id,
						name: data.cs[i].n,
						child: [],
						_folder: data.cs[i].ifr,
						_name: data.cs[i].n,
						_show: true,
						_active: false,
						_ajax: false,
						_loading: false,
						_expend: false,
						$parent: null, // 父级对象
						$index: i
					});
				}

				vm.tree = tree;
			});
		}

		$ctrl.$onRendered = function () {
			// 左侧整体右键
			$.contextMenu({
				selector: '#game-list-doc .menu',
				callback: function (key, options) {
					vm.add.$docId = vm.$categoryId;
					vm.add.$parentId = vm.$categoryId;
					switch (key) {
					case 'addFolder':
						vm.$mouseObj = {};
						// 追加树之后
						$('#game-list-doc #tree').after($("#game-list-doc .menu .add"));

						vm.add.type = 'folder';
						vm.add.show = true;

						// 获取焦点
						$('#game-list-doc .menu #add-title').focus();
						break;
					case 'addFile':
						vm.$mouseObj = {};
						// 追加树之后
						$('#game-list-doc #tree').after($("#game-list-doc .menu .add"));

						vm.add.type = 'file';
						vm.add.show = true;

						// 获取焦点
						$('#game-list-doc .menu #add-title').focus();
						break;
					}
				},
				items: {
					"addFolder": {
						name: "新建文件夹",
						disabled: false
					},
					"addFile": {
						name: "新建文件",
						disabled: false
					}
				}
			});

			// 文件夹上右键
			$.contextMenu({
				selector: '#game-list-doc .menu .folder',
				callback: function (key, options) {
					vm.add.$docId = $(this).attr('doc-id');
					vm.add.$parentId = $(this).attr('parent-id') || vm.$categoryId;
					switch (key) {
					case 'addFolder':
						// 追加树之后
						$(this).after($("#game-list-doc .menu .add"));

						vm.add.type = 'folder';
						vm.add.show = true;

						// 获取焦点
						$('#game-list-doc .menu #add-title').focus();
						break;
					case 'addFile':
						// 追加树之后
						$(this).after($("#game-list-doc .menu .add"));

						vm.add.type = 'file';
						vm.add.show = true;

						// 获取焦点
						$('#game-list-doc .menu #add-title').focus();
						break;
					case 'delete':
						// 查询父级
						var parentChilds, parentId;
						if (vm.$mouseObj.$parent == null) { // 父级为根元素
							parentChilds = vm.tree;
							parentId = vm.$treeId;
						} else {
							parentChilds = vm.$mouseObj.$parent.child;
							parentId = vm.$mouseObj.$parent.id;
						}

						// 查找位于父级的index
						var index = 0;
						for (var i = 0, j = parentChilds.size(); i < j; i++) {
							if (parentChilds[i].id == vm.$mouseObj.id) {
								index = i;
							}
						}

						post('/api/doc/deleteFolder', {
							parent: parentId,
							index: index
						}, null, '删除失败：', function () {
							parentChilds.removeAt(index);
						});
						break;
					}
				},
				items: {
					"addFolder": {
						name: "新建文件夹",
						disabled: false
					},
					"addFile": {
						name: "新建文件",
						disabled: false
					},
					"delete": {
						name: "删除",
						disabled: false
					}
				}
			});

			// 文件上右键
			$.contextMenu({
				selector: '#game-list-doc .menu .file',
				callback: function (key, options) {
					switch (key) {
					case 'delete':
						// 查询父级
						var parentChilds, parentId;

						if (vm.$mouseObj.$parent == null) { // 父级为根元素
							parentChilds = vm.tree;
							parentId = vm.$treeId;
						} else {
							parentChilds = vm.$mouseObj.$parent.child;
							parentId = vm.$mouseObj.$parent.id;
						}

						// 查找位于父级的index
						var index = 0;
						for (var i = 0, j = parentChilds.size(); i < j; i++) {
							if (parentChilds[i].id == vm.$mouseObj.id) {
								index = i;
							}
						}

						post('/api/game/operate', {
							id: vm.$mouseObj.id,
							type: 'delete',
							docParent: parentId,
							docIndex: index
						}, null, '删除失败：', function () {
							parentChilds.removeAt(index);

							// 如果删除是当前文章，返回分类首页
							if (vm.$mouseObj.id == vm.nowTopicId) {
								avalon.router.navigate('/g/' + avalon.vmodels.gameBase.game.Id + '/' + avalon.vmodels.gameBase.category);
							}
						});
						break;
					}
				},
				items: {
					"delete": {
						name: "删除",
						disabled: false
					}
				}
			});

			// 左侧菜单初始化拖拽事件
			vm.$sortable();
		}
	});
});