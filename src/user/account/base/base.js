'use strict';

define("userAccountBase", ['jquery', 'jquery.timeago'], function ($) {
	var vm = avalon.define({
		$id: "userAccountBase"
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {
			//jbox插件
			jbox();

			// timeago 插件
			$('#user-account-base .timeago').timeago();
		}
	});
});