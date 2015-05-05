'use strict';

define("userYuqingAnalyse", ['jquery', 'js/public/avalon.table'], function ($) {
	var types = {
		0: '形容词',
		1: '名词',
		2: '动词',
		3: '程度副词',
		4: '否定词',
		5: '介词',
		6: '语气助词'
	}

	var vm = avalon.define({
		$id: "userYuqingAnalyse",
		$opts: {
			title: {
				'_id': {
					name: '词语',
					sort: true,
					add: true,
					search: true
				},
				't': {
					name: '类型',
					add: true,
					mutiple: types
				},
				'p': {
					name: '程度（越大越积极，反之消极）',
					add: true,
					select: {
						'-5': -5,
						'-4': -4,
						'-3': -3,
						'-2': -2,
						'-1': -1,
						'0': 0,
						'1': 1,
						'2': 2,
						'3': 3,
						'4': 4,
						'5': 5
					}
				},
				'i': {
					name: '屏蔽情感倾向的词性',
					add: true,
					mutiple: types
				}
			},
			add: true,
			_delete: true,
			update: true,
			like: true,
			url: '/api/yuqing/analyse'
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {

		}
	});
});