'use strict';

define("gameAdmin", ['jquery'], function ($) {

	var vm = avalon.define({
		$id: "gameAdmin",
		info: ''
	});

	return avalon.controller(function ($ctrl) {

		$ctrl.$onEnter = function (param, rs, rj) {
			// 赋值base当前所在位置
			avalon.vmodels.gameBase.baseRouter = 'admin';
		}

		$ctrl.$onRendered = function () {}

	});

});