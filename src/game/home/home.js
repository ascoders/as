'use strict';

define("gameHome", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "gameHome",
		images: [], //应用截图
		imagePage: 1, //当前截图页
		changeImagePage: function (page) { //改变截图页
			vm.imagePage = page;
			var standardWidth = $("#game-base #j-main-container").width() - 40;
			$("#over-con-ul").animate({
				marginLeft: -(standardWidth / 3 * 2 + 80) * (page - 1)
			}, 300);
		}
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			avalon.vmodels.gameBase.baseRouter = '';

			// 改变页面titile
			document.title = avalon.vmodels.gameBase.game.Name;
		}
		$ctrl.$onRendered = function () {
			//设置应用截图框架高度
			var standardWidth = $("#game-base #j-main-container").width() - 40;

			var imageHeight = standardWidth / 3 * 1.6;
			$("#game-home .over-con").css("height", imageHeight + "px").css("width", imageHeight / 1.6 * 2 + 80);

			//设置图片宽高
			$("#game-home .introduce_image").css("width", imageHeight / 1.6).css("height", imageHeight);

			//设置左右按钮位置
			$("#game-home .image_hover").css("width", standardWidth / 8 + "px").css("height", standardWidth / 3 * 1.6 + "px").css("line-height", standardWidth / 3 * 1.6 + "px");
		}
	});
});