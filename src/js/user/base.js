'use strict';

define("userBase", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "userBase",
		lists: [{ //左侧菜单
			title: '我的账号',
			url: 'account',
			icon: 'fa-user',
			show: true,
			ieVersion: ieVersion(),
			childs: [{
				name: '基本信息',
				icon: 'fa-bar-chart-o',
				url: 'base'
			}, {
				name: '修改密码',
				icon: 'fa-key',
				url: 'password'
			}, {
				name: '消息中心',
				icon: 'fa-pagelines',
				url: 'message'
			}, {
				name: '账户充值',
				icon: 'fa-skype',
				url: 'recharge'
			}, {
				name: '充值记录',
				icon: 'fa-history',
				url: 'history'
			}, {
				name: '绑定邮箱',
				icon: 'fa-envelope',
				url: 'email'
			}, {
				name: '第三方平台绑定',
				icon: 'fa-refresh',
				url: 'oauth'
			}]
		}, {
			title: 'Erp系统',
			url: 'erp',
			icon: 'fa-street-view',
			show: true,
			childs: [{
				name: '领取工资',
				icon: 'fa-usd',
				url: 'salary'
			}]
		}, {
			title: '网站管理', // admin可见
			url: 'admin',
			icon: 'fa-cog',
			show: false,
			childs: [{
				name: '账号管理',
				icon: 'fa-user',
				url: 'member'
			}, {
				name: '职位管理',
				icon: 'fa-usd',
				url: 'job'
			}]
		}, {
			title: '舆情分析',
			url: 'yuqing',
			icon: 'fa-sellsy',
			show: false,
			childs: [{
				name: '分词词库',
				icon: 'fa-file-word-o',
				url: 'split'
			}, {
				name: '分析词库',
				icon: 'fa-file-archive-o',
				url: 'analyse'
			}, {
				name: '管理',
				icon: 'fa-cogs',
				url: 'operate'
			}]
		}],
		title: '', //当前标题
		category: '', // 当前分类
		page: '', // 当前二级分类
		toggleShow: function () { // 显示子元素
			var url = $(this).attr('name')
			$('#user-base .j-body').each(function () {
				if (url != $(this).attr('name')) {
					$(this).slideUp(200)
				} else {
					$(this).slideToggle(200)
				}
			});
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			$.when(avalon.vmodels.global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
				if (!global.myLogin) {
					avalon.router.navigate('/');
					return;
				}

				for (var key in vm.lists.$model) {
					if (!vm.lists.$model.hasOwnProperty(key) || key === "hasOwnProperty") {
						continue;
					}

					if (vm.lists[key].show) { // 默认显示则跳过
						continue
					}

					// 管理员才显示的
					if (vm.lists[key].url == 'admin' && avalon.vmodels.global.my.type == 0) { // 超级管理员
						vm.lists[key].show = true
						continue
					}

					// 拥有权限才显示的
					if ($.inArray(vm.lists[key].url, avalon.vmodels.global.my.power) > -1) {
						vm.lists[key].show = true
					}
				}
			});
		}

		$ctrl.$onRendered = function () {

		}
	});
});