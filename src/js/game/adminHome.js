'use strict';

define("gameAdminHome", ['jquery'], function ($) {

	var vm = avalon.define({
		$id: "gameAdminHome"
	});

	return avalon.controller(function ($ctrl) {

		$ctrl.$onEnter = function (param, rs, rj) {
			// 设置状态
			avalon.vmodels.gameAdmin.info = '';

			// 改变页面titile
			document.title = '总览 - 管理 - ' + avalon.vmodels.gameBase.game.Name;
		}
		$ctrl.$onRendered = function () {}

	});

});