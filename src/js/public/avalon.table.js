'use strict';

define(["jquery", "text!html/public/avalon.table.html", 'avalon.page', 'jquery.timeago', 'jquery.contextMenu', 'jquery.jbox'], function ($, tableHtml, pageVm) {

    var uiName = 'table';

    avalon.ui[uiName] = function (element, data, vmodels) {

        var innerHTML = element.innerHTML;
        avalon.clearHTML(element);

        // 参数
        var opts = data[uiName + "Options"];

        var vm = avalon.define({
            $id: data[uiName + "Id"],
            title: {},
            content: [], //内容
            sort: '', //排序对象
            pagin: '', // 分页
            searchCondition: '', //搜索条件字段
            search: '', // 搜索关键字
            checkboxs: {}, //当前多选项
            checkText: [], //多选项内容
            add: false, // 是否可添加
            searchMethod: 'accuracy', // 搜索方式
            $init: function () { // 初始化
                // 执行opts的初始化
                if (typeof opts.onInit === "function") {
                    opts.onInit(opts);
                }

                // 不同state或param参数切换时，分页按钮隐藏
                if (mmState.prevState !== null && mmState.prevState.stateName !== mmState.currentState.stateName) {
                    vm.pagin = ''
                } else if (mmState.prevState !== null && mmState.prevState.path !== mmState.currentState.path) {
                    vm.pagin = ''
                }

                // 设置默认排序字段
                for (var key in opts.title) {
                    if (opts.title[key].sort) {
                        opts.sort = key;
                        break;
                    }
                }

                //为title增加_add参数，与新增输入框绑定 并赋值其他参数
                for (var key in opts.title) {
                    opts.title[key]._add = '';

                    opts.title[key]._type = 'text';

                    if (!$.isEmptyObject(opts.title[key].select)) {
                        opts.title[key]._type = 'select';
                    }

                    if (!$.isEmptyObject(opts.title[key].mutiple)) {
                        opts.title[key]._type = 'mutiple';
                        opts.title[key]._add = [];
                        opts.title[key]._html = '';
                    }

                    if (opts.title[key].search) {
                        vm.searchCondition = key;
                    }
                }

                // 将参数（opts）赋值到vm变量中
                vm.title = opts.title;
                vm.sort = opts.sort;
                vm.add = opts.add;

                // 初始化分页插件
                pageVm.$from = 0;

                pageVm.$post = function () {
                    post(opts.url, {
                        type: 'get',
                        from: pageVm.$from,
                        number: pageVm.$number,
                        sort: vm.sort,
                        like: vm.search,
                        likeKey: vm.searchCondition,
                        likeMethod: vm.searchMethod
                    }, null, '获取信息失败：', function (data) {
                        // 刷新content
                        var _array = new Array();
                        for (var key = 0, k = data.lists.length; key < k; key++) {

                            var obj = {};

                            for (var __key in opts.title) {
                                if (!opts.title.hasOwnProperty(__key) || __key === "hasOwnProperty") {
                                    continue;
                                }


                                obj[__key] = {};
                                obj[__key].val = data.lists[key][__key]; //真实值

                                //文本类型
                                obj[__key].type = 'text';

                                // 多选框类型 数组类型不能直接toString
                                if (!$.isEmptyObject(opts.title[__key].mutiple)) {
                                    obj[__key].type = 'mutiple';

                                    if (data.lists[key][__key] != null && data.lists[key][__key].length > 0) {
                                        var strArray = [];

                                        // val清空
                                        obj[__key].val = [];

                                        for (var ___key in data.lists[key][__key]) {
                                            if (!data.lists[key][__key].hasOwnProperty(___key) || ___key === "hasOwnProperty") {
                                                continue;
                                            }

                                            // 如果值不在title mutiple数组中，值为空
                                            if (opts.title[__key].mutiple[data.lists[key][__key][___key]] == undefined) {
                                                continue;
                                            }

                                            obj[__key].val.push(data.lists[key][__key][___key]);

                                            strArray.push(opts.title[__key].mutiple[data.lists[key][__key][___key]]);
                                        }

                                        obj[__key].html = strArray.join(',');
                                    } else {
                                        obj[__key].html = '';
                                        obj[__key].val = [];
                                    }
                                } else {
                                    obj[__key].html = data.lists[key][__key].toString(); //显示值

                                    // 搜索模式下，高亮
                                    if (vm.search != '' && vm.searchCondition != '') {
                                        obj[__key].html = obj[__key].html.replace(eval('/' + vm.search + '/g'), '<span class="search-key">' + vm.search + '</span>');
                                    }
                                }

                                //时间类型
                                if (opts.title[__key].time) {
                                    obj[__key].type = 'time';
                                }

                                //下拉选框类型
                                if (!$.isEmptyObject(opts.title[__key].select)) {
                                    obj[__key].type = 'select';
                                    obj[__key].select = opts.title[__key].select;
                                }


                            }
                            _array.push(obj);
                        }

                        // 赋值给内容
                        vm.content = _array;

                        // 刷新分页模块
                        pageVm.$fresh(data.count);
                    });
                }

                // 获取当前页面信息
                pageVm.$post();

                // 填充模版
                element.innerHTML = tableHtml + pageVm.$page;

                avalon.scan(element, [vm].concat(vmodels));

                if (typeof vm.onInit === "function") {
                    vm.onInit.call(element, vm, options, vmodels);
                }
            },
            $remove: function () { // 销毁
                element.innerHTML = ""
            },
            rendered: function (type) { // 渲染完毕
                if (type != 'add') {
                    return
                }

                $('.timeago').timeago()

                //右键菜单
                $.contextMenu('destroy')
                $.contextMenu({
                    selector: '.avalon-table-tr',
                    callback: function (key) {
                        // 操作第几行
                        var index = $(this).attr('array-index');

                        // 操作类别
                        switch (key) {
                        case 'delete':
                            var key = vm.title[vm.sort.replace('-', '')].name;
                            var value = vm.content[index][vm.sort.replace('-', '')].val;

                            confirm(key + '：' + value + ' 确认删除吗', function () {
                                post(opts.url, {
                                    type: 'delete',
                                    _id: vm.content[index]['_id'].val
                                }, '删除成功', '删除失败：', function () { //删除成功
                                    // 如果页面记录数大于1或者是第一页，刷新本页面
                                    if (vm.content.size() > 1 || vm.$temp.from == 0) {
                                        pageVm.jump(pageVm.page, true);
                                    } else { // 刷新到上一页
                                        pageVm.jump(pageVm.page - 1, true);
                                    }
                                });
                            });
                            break;
                        }
                    },
                    items: {
                        "delete": {
                            name: "删除",
                            disabled: !opts._delete
                        }
                    }
                });
            },
            modalRendered: function () { // 模态框渲染完毕，初始化模态框，仅一次
                if ($.isEmptyObject(opts.addModal)) {
                    opts.addModal = new jBox('Confirm', {
                        minWidth: '200px',
                        content: $('#modal'),
                        animation: 'flip',
                        confirmButton: '确定',
                        cancelButton: '取消',
                        overlay: false,
                        zIndex: 10000,
                        confirm: function () {
                            var params = {
                                type: 'add'
                            };

                            for (var key in vm.title.$model) {
                                if (typeof (vm.title[key]._add) == 'string') {
                                    params[key] = vm.title[key]._add;
                                } else {
                                    params[key] = vm.title[key]._add.$model;
                                }

                            }

                            //发送请求
                            post(opts.url, params, '新增成功', '新增失败：', function () { //新增成功
                                // 刷新本页
                                pageVm.jump(pageVm.page, true);
                            });
                        }
                    });
                }

                if ($.isEmptyObject(opts.checkboxModal)) {
                    opts.checkboxModal = new jBox('Modal', {
                        minWidth: '200px',
                        content: $('#modal-checkbox'),
                        animation: 'flip',
                        overlay: false,
                        zIndex: 10001
                    });
                }
            },
            addClick: function () { // 新增按钮点击，弹出窗
                // 先关闭多选框
                opts.checkboxModal.close();

                opts.addModal.open();
            },
            changeMutiple: function (index, key) { // 【更新】打开多选框
                // 更新枚举列表
                vm.checkboxs = vm.title[key].mutiple;

                // 更新当前数组
                vm.checkText = vm.content[index][key].val;

                // 转化为字符串（Number类型无法绑定）
                for (var i = 0, j = vm.checkText.$model.length; i < j; i++) {
                    vm.checkText[i] = vm.checkText[i].toString()
                }

                console.log(vm.checkboxs.$model, vm.checkText.$model)

                // 保存此时checkText 值拷贝
                var tempCheckText = vm.checkText.$model.slice(0);

                // 关闭后保存
                opts.checkboxModal.options.onClose = function () {
                    // 如果没有更改，直接退出
                    if (vm.checkText.$model.sort().toString() == tempCheckText.sort().toString()) {
                        return;
                    }

                    var params = {
                        type: 'update',
                        _id: vm.content[index]['_id'].val
                    };

                    // 增加修改元素
                    params[key] = vm.checkText.$model;

                    post(opts.url, params, '更新成功', '更新失败：', function () { // 更新成功
                        // 更新数据
                        vm.content[index][key].val = vm.checkText;

                        var strArray = [];
                        for (var _key in vm.checkText.$model) {
                            strArray.push(vm.title[key].mutiple[vm.checkText[_key]]);
                        }

                        vm.content[index][key].html = strArray.join(',');
                    });
                }

                opts.checkboxModal.open();
            },
            addMutiple: function (key) { // 【新增】打开多选框
                // 更新枚举列表
                vm.checkboxs = vm.title[key].mutiple;

                // 更新当前数组
                vm.checkText = vm.title[key]._add;

                // 模态框打开
                opts.checkboxModal.open();

                // 关闭后赋值
                opts.checkboxModal.options.onClose = function () {
                    // 赋值
                    vm.title[key]._add = vm.checkText.$model;

                    var _array = [];
                    for (var _key in vm.title[key]._add.$model) {
                        if (vm.title[key].mutiple[vm.title[key]._add.$model[_key]] != undefined) {
                            _array.push(vm.title[key].mutiple[vm.title[key]._add.$model[_key]]);
                        }
                    }
                    vm.title[key]._html = _array.join(',');
                }
            },
            $searchTimeout: null,
            changeSearch: function (value) { //改变搜索内容
                if (vm.$searchTimeout != null) {
                    clearTimeout(vm.$searchTimeout)
                }

                vm.$searchTimeout = setTimeout(function () {
                    vm.like = value
                    pageVm.jump(1, true)
                }, 200)
            },
            searchMethodChange: function () { // 改变搜索方式
                vm.$searchTimeout = setTimeout(function () {
                    pageVm.jump(1, true)
                }, 200)
            },
            changeSort: function (sort) { //改变排序
                if (sort == vm.sort) { // 同样的点击，取反
                    sort = '-' + sort;
                }

                vm.sort = sort;

                pageVm.jump(pageVm.page, true);
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
                    _id: vm.content[index]['_id'].val
                };

                //增加修改元素
                params[key] = vm.content[index][key].val;

                post(opts.url, params, '更新成功', '更新失败：', null, null);
            },
            update: function (index, key) { // 双击修改
                //只有text类型才能双击修改
                if (vm.content[index][key].type != 'text') {
                    return;
                }

                //暂存数据
                opts.text = vm.content[index][key].val;

                $(this).find('span').hide();
                $(this).addClass('f-p0').find('input').show().focus();
            },
            onBlur: function (index, key) { // 输入框取消焦点，更新
                //如果值没变，不做处理
                if (vm.content[index][key].val == opts.text) {
                    $(this).hide();
                    $(this).parent().parent().removeClass('f-p0').find('span').show();
                    return;
                }

                var _this = this;

                //状态不可编辑
                $(_this).attr('disabled', 'disabled');

                var params = {
                    type: 'update',
                    _id: vm.content[index]['_id'].val
                };

                //增加修改元素
                params[key] = vm.content[index][key].val;

                post(opts.url, params, '更新成功', '更新失败：', function () { //更新成功
                    //修改元素
                    vm.content[index][key].html = vm.content[index][key].val;
                    //输入框消失，还原文字
                    $(_this).hide();
                    $(_this).parent().parent().removeClass('f-p0').find('span').show();
                    $(_this).removeAttr('disabled');
                }, function () { // 更新失败，该字段回退到原来的值
                    //元素内容还原成未修改之前
                    vm.content[index][key].html = vm.content[index][key].val = vm.$temp.text;
                    //输入框消失，还原文字
                    $(_this).hide();
                    $(_this).parent().parent().removeClass('f-p0').find('span').show();
                    $(_this).removeAttr('disabled');
                });
            },
            $skipArray: ['onChange', '$temp']
        });

        return vm;
    }

    avalon.ui[uiName].defaults = {
        add: false,
        _delete: false,
        update: false,
        like: false,
        title: {},
        url: '',
        sort: '',
        from: 0,
        number: 10,
        text: '', //暂存修改前的文本框值
        addModal: {}, //新增模态框
        checkboxModal: {} //多选框模态框
    }

    return avalon;
});