'use strict';

define("userAdminMember", ['jquery', 'js/public/avalon.table'], function ($) {
	var vm = avalon.define({
		$id: "userAdminMember",
		$opts: {
			title: {
				'_id': {
					name: 'ID',
					sort: true
				},
				'n': {
					name: '昵称',
					add: true,
					search: true
				},
				'e': {
					name: '邮箱',
					add: true
				},
				'mo': {
					name: '余额',
					add: true
				},
				'l': {
					name: '登录次数'
				},
				'la': {
					name: '最后登录时间',
					time: true
				},
				'po': {
					name: '后台权限',
					add: true,
					mutiple: {
						'backingdoms': '回到三国',
						'threeRoad': '三途路',
						'yuqing': '舆情分析'
					}
				},
				't': {
					name: '帐号类型',
					add: true
				}
			},
			url: '/api/user/member',
			add: false,
			_delete: false,
			update: true,
			like: true,
			from: 0,
			onInit: function (opts) {
				// 获取type
				post('/api/user/jobs', {}, null, '职位列表获取失败：', function (data) {
					//账号类型select
					var select = {};
					for (var key in data) {
						select[data[key]._id] = data[key].n;
					}

					opts.title['t'].select = select;
				});
			}
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {

		}
	});
});