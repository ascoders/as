'use strict';

define("userAdminJob", ['jquery', 'js/public/avalon.table'], function ($) {
	var vm = avalon.define({
		$id: "userAdminJob",
		$opts: {
			title: {
				'_id': {
					name: 'ID',
					sort: true,
					add: true
				},
				'n': {
					name: '职位名称',
					add: true,
					search: true
				},
				'd': {
					name: '工作职责',
					add: true
				},
				'r': {
					name: '岗位要求',
					add: true
				},
				's': {
					name: '月薪',
					add: true
				},
				'u': {
					name: '每日上传限额',
					add: true
				}
			},
			add: true,
			_delete: true,
			update: true,
			like: true,
			url: '/api/user/job'
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {

		}
	});
});