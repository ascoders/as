'use strict';

define("tempTable", ['jquery', 'jquery.timeago', 'jquery.contextMenu', 'jquery.jbox'], function ($) {
	return avalon.define({
		$id: "tempTable",
		title: {}, //表题
		content: [], //内容
		sort: '', //排序对象
		pagin: '',
		searchCondition: '', //搜索条件字段
		checkboxs: {}, //当前多选项
		checkText: [], //多选项内容
		temp: {
			url: '',
			from: 0,
			number: 0,
			text: '', //暂存修改前的文本框值
			addModal: {}, //新增模态框
			checkboxModal: {}, //多选框模态框
			delete: false, //是否允许删除
		},
		rendered: function (type) { // 渲染完毕
			$('.timeago').timeago();

			//右键菜单
			$.contextMenu({
				selector: '.temp-table-tr',
				callback: function (key) {
					// 操作第几行
					var index = $(this).attr('array-index');

					// 操作类别
					switch (key) {
					case 'delete':
						var key = avalon.vmodels.tempTable.title[avalon.vmodels.tempTable.sort.replace('-', '')].name;
						var value = avalon.vmodels.tempTable.content[index][avalon.vmodels.tempTable.sort.replace('-', '')].val;

						confirm(key + '：' + value + ' 确认删除吗', function () {
							post(avalon.vmodels.tempTable.temp.url, {
								type: 'delete',
								_id: avalon.vmodels.tempTable.content[index]['_id'].val,
							}, '删除成功', '删除失败：', function () { //删除成功
								// 如果页面记录数大于1或者是第一页，刷新本页面
								if (avalon.vmodels.tempTable.content.size() > 1 || avalon.vmodels.tempTable.temp.from == 0) {
									avalon.router.navigate(window.location.pathname + '?from=' + avalon.vmodels.tempTable.temp.from + '&number=' + avalon.vmodels.tempTable.temp.number + '&sort=' + avalon.vmodels.tempTable.sort, {
										reload: true,
									});
								} else { // 刷新到上一页
									avalon.router.navigate(window.location.pathname + '?from=' + (avalon.vmodels.tempTable.temp.from - avalon.vmodels.tempTable.temp.number) + '&number=' + avalon.vmodels.tempTable.temp.number + '&sort=' + avalon.vmodels.tempTable.sort);
								}
							});
						});
						break;
					}
				},
				items: {
					"delete": {
						name: "删除",
						disabled: !avalon.vmodels.tempTable.temp.delete,
					},
				}
			});
		},
		modalRendered: function () { // 模态框渲染完毕，初始化模态框，仅一次
			if ($.isEmptyObject(avalon.vmodels.tempTable.temp.addModal)) {
				avalon.vmodels.tempTable.temp.addModal = new jBox('Confirm', {
					minWidth: '200px',
					content: $('#modal'),
					animation: 'flip',
					confirmButton: '确定',
					cancelButton: '取消',
					overlay: false,
					zIndex: 10000,
					confirm: function () {
						var params = {
							type: 'add',
						};

						for (var key in avalon.vmodels.tempTable.title.$model) {
							if (typeof (avalon.vmodels.tempTable.title[key]._add) == 'string') {
								params[key] = avalon.vmodels.tempTable.title[key]._add;
							} else {
								params[key] = avalon.vmodels.tempTable.title[key]._add.$model;
							}

						}

						//发送请求
						post(avalon.vmodels.tempTable.temp.url, params, '新增成功', '新增失败：', function () { //新增成功
							// 刷新本页
							avalon.router.navigate(window.location.pathname + '?from=' + avalon.vmodels.tempTable.temp.from + '&number=' + avalon.vmodels.tempTable.temp.number + '&sort=' + avalon.vmodels.tempTable.sort, {
								reload: true,
							});
						});
					}
				});
			}

			if ($.isEmptyObject(avalon.vmodels.tempTable.temp.checkboxModal)) {
				avalon.vmodels.tempTable.temp.checkboxModal = new jBox('Modal', {
					minWidth: '200px',
					content: $('#modal-checkbox'),
					animation: 'flip',
					overlay: false,
					zIndex: 10001,
				});
			}
		},
		add: function () { // 新增按钮点击，弹出窗
			// 先关闭多选框
			avalon.vmodels.tempTable.temp.checkboxModal.close();

			avalon.vmodels.tempTable.temp.addModal.open();
		},
		changeMutiple: function (index, key) { // 【更新】打开多选框
			// 更新枚举列表
			avalon.vmodels.tempTable.checkboxs = avalon.vmodels.tempTable.title[key].mutiple;

			// 更新当前数组
			avalon.vmodels.tempTable.checkText = avalon.vmodels.tempTable.content[index][key].val;

			// 保存此时checkText 值拷贝
			var tempCheckText = avalon.vmodels.tempTable.checkText.$model.slice(0);

			// 关闭后保存
			avalon.vmodels.tempTable.temp.checkboxModal.options.onClose = function () {
				// 如果没有更改，直接退出
				if (avalon.vmodels.tempTable.checkText.$model.sort().toString() == tempCheckText.sort().toString()) {
					return;
				}

				var params = {
					type: 'update',
					_id: avalon.vmodels.tempTable.content[index]['_id'].val,
				};

				// 增加修改元素
				params[key] = avalon.vmodels.tempTable.checkText.$model;

				post(avalon.vmodels.tempTable.temp.url, params, '更新成功', '更新失败：', function () { // 更新成功
					// 更新数据
					avalon.vmodels.tempTable.content[index][key].val = avalon.vmodels.tempTable.checkText;

					var strArray = [];
					for (var _key in avalon.vmodels.tempTable.checkText.$model) {
						strArray.push(avalon.vmodels.tempTable.title[key].mutiple[avalon.vmodels.tempTable.checkText[_key]]);
					}

					avalon.vmodels.tempTable.content[index][key].html = strArray.join(',');
				});
			}

			avalon.vmodels.tempTable.temp.checkboxModal.open();
		},
		addMutiple: function (key) { // 【新增】打开多选框
			// 更新枚举列表
			avalon.vmodels.tempTable.checkboxs = avalon.vmodels.tempTable.title[key].mutiple;

			// 更新当前数组
			avalon.vmodels.tempTable.checkText = avalon.vmodels.tempTable.title[key]._add;

			// 模态框打开
			avalon.vmodels.tempTable.temp.checkboxModal.open();

			// 关闭后赋值
			avalon.vmodels.tempTable.temp.checkboxModal.options.onClose = function () {
				// 赋值
				avalon.vmodels.tempTable.title[key]._add = avalon.vmodels.tempTable.checkText.$model;

				var _array = [];
				for (var _key in avalon.vmodels.tempTable.title[key]._add.$model) {
					if (avalon.vmodels.tempTable.title[key].mutiple[avalon.vmodels.tempTable.title[key]._add.$model[_key]] != undefined) {
						_array.push(avalon.vmodels.tempTable.title[key].mutiple[avalon.vmodels.tempTable.title[key]._add.$model[_key]]);
					}
				}
				avalon.vmodels.tempTable.title[key]._html = _array.join(',');
			}
		},
		onChange: function (params) {
			//额外url参数
			var opts = {};

			//排序参数
			opts.sort = params.query.sort;

			if (opts.sort == '' || opts.sort == undefined) {
				for (var key in params.title) {
					if (params.title[key].sort) {
						opts.sort = key;
						break;
					}
				}
			}

			//获取区间
			params.query.from = params.query.from || 0;
			params.query.number = params.query.number || 10;

			//为title增加_add参数，与新增输入框绑定
			for (var key in params.title) {
				params.title[key]._add = '';

				params.title[key]._type = 'text';

				if (!$.isEmptyObject(params.title[key].select)) {
					params.title[key]._type = 'select';
				}

				if (!$.isEmptyObject(params.title[key].mutiple)) {
					params.title[key]._type = 'mutiple';
					params.title[key]._add = [];
					params.title[key]._html = '';
				}

				if (params.title[key].search) {
					avalon.vmodels.tempTable.searchCondition = key;
				}
			}

			//保存数据
			avalon.vmodels.tempTable.temp.from = params.query.from;
			avalon.vmodels.tempTable.temp.number = params.query.number;
			avalon.vmodels.tempTable.temp.url = params.url;
			avalon.vmodels.tempTable.temp.delete = params.delete;
			avalon.vmodels.tempTable.title = params.title;
			avalon.vmodels.tempTable.sort = opts.sort;

			post(params.url, {
				type: 'get',
				from: params.query.from,
				number: params.query.number,
				sort: opts.sort,
			}, null, '获取信息失败：', function (data) {
				// 刷新content
				var _array = new Array();
				for (var key in data.lists) {
					var obj = {};
					for (var __key in params.title) {
						obj[__key] = {};
						obj[__key].val = data.lists[key][__key]; //真实值

						//文本类型
						obj[__key].type = 'text';

						// 多选框类型 数组类型不能直接toString
						if (!$.isEmptyObject(params.title[__key].mutiple)) {
							obj[__key].type = 'mutiple';

							if (data.lists[key][__key] != null && data.lists[key][__key].length > 0) {
								var strArray = [];

								// val清空
								obj[__key].val = [];

								for (var ___key in data.lists[key][__key]) {
									// 如果值不在title mutiple数组中，值为空
									if (params.title[__key].mutiple[data.lists[key][__key][___key]] == undefined) {
										continue;
									}

									obj[__key].val.push(data.lists[key][__key][___key]);

									strArray.push(params.title[__key].mutiple[data.lists[key][__key][___key]]);
								}

								obj[__key].html = strArray.join(',');
							} else {
								obj[__key].html = '';
								obj[__key].val = [];
							}
						} else {
							obj[__key].html = data.lists[key][__key].toString(); //显示值
						}

						//时间类型
						if (params.title[__key].time) {
							obj[__key].type = 'time';
						}

						//下拉选框类型
						if (!$.isEmptyObject(params.title[__key].select)) {
							obj[__key].type = 'select';
							obj[__key].select = params.title[__key].select;
						}


					}
					_array.push(obj);
				}
				avalon.vmodels.tempTable.content = _array;

				//生成分页
				avalon.vmodels.tempTable.pagin = createPagin(params.query.from, params.query.number, data.count, opts);
			});
		},
		changeSearch: function () { //改变搜索内容
			console.log($(this).val());
		},
		changeSort: function (sort) { //改变排序
			if (sort == avalon.vmodels.tempTable.sort) { // 同样的点击，取反
				sort = '-' + sort;
			}
			avalon.router.navigate(window.location.pathname + '?from=' + avalon.vmodels.tempTable.temp.from + '&number=' + avalon.vmodels.tempTable.temp.number + '&sort=' + sort);
		},
		compareSort: function (value1, value2) { // 显示用，判断是否为该sort
			if (value1.replace('-', '') == value2.replace('-', '')) {
				return true;
			} else {
				return false;
			}
		},
		asc: function (value) { // 显示用，正序还是逆序
			if (value.indexOf('-') > -1) {
				return false;
			} else {
				return true;
			}
		},
		selectChange: function (index, key) { //下拉选项修改
			var params = {
				type: 'update',
				_id: avalon.vmodels.tempTable.content[index]['_id'].val,
			};

			//增加修改元素
			params[key] = avalon.vmodels.tempTable.content[index][key].val;

			post(avalon.vmodels.tempTable.temp.url, params, '更新成功', '更新失败：', null, null);
		},
		update: function (index, key) { // 双击修改
			//只有text类型才能双击修改
			if (avalon.vmodels.tempTable.content[index][key].type != 'text') {
				return;
			}

			//暂存数据
			avalon.vmodels.tempTable.temp.text = avalon.vmodels.tempTable.content[index][key].val;

			$(this).find('span').hide();
			$(this).addClass('f-p0').find('input').show().focus();
		},
		onBlur: function (index, key) { // 输入框取消焦点，更新
			//如果值没变，不做处理
			if (avalon.vmodels.tempTable.content[index][key].val == avalon.vmodels.tempTable.temp.text) {
				$(this).hide();
				$(this).parent().parent().removeClass('f-p0').find('span').show();
				return;
			}

			var _this = this;

			//状态不可编辑
			$(_this).attr('disabled', 'disabled');

			var params = {
				type: 'update',
				_id: avalon.vmodels.tempTable.content[index]['_id'].val,
			};

			//增加修改元素
			params[key] = avalon.vmodels.tempTable.content[index][key].val;

			post(avalon.vmodels.tempTable.temp.url, params, '更新成功', '更新失败：', function () { //更新成功
				//修改元素
				avalon.vmodels.tempTable.content[index][key].html = avalon.vmodels.tempTable.content[index][key].val;
				//输入框消失，还原文字
				$(_this).hide();
				$(_this).parent().parent().removeClass('f-p0').find('span').show();
				$(_this).removeAttr('disabled');
			}, function () { // 更新失败，该字段回退到原来的值
				//元素内容还原成未修改之前
				avalon.vmodels.tempTable.content[index][key].html = avalon.vmodels.tempTable.content[index][key].val = avalon.vmodels.tempTable.temp.text;
				//输入框消失，还原文字
				$(_this).hide();
				$(_this).parent().parent().removeClass('f-p0').find('span').show();
				$(_this).removeAttr('disabled');
			});
		},
		$skipArray: ['onChange', 'temp'],
	});
});