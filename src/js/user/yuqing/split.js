'use strict';

define("userYuqingSplit", ['jquery', 'avalon.table'], function ($) {
	var vm = avalon.define({
		$id: "userYuqingSplit",
		$opts: {
			title: {
				'_id': {
					name: '分词',
					sort: true,
					add: true,
					search: true
				},
				'h': {
					name: '频率',
					add: true
				},
				't': {
					name: '词性',
					add: true
				}
			},
			add: true,
			_delete: true,
			update: true,
			like: true,
			url: '/api/yuqing/split'
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {

		}
	});
});