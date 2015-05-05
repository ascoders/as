'use strict';
define("gameGame", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "gameGame",
		category: '', //当前分类
		lists: [], //列表页
		pagin: '', //分页html
		type: '' //当前类型
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$vmodels = [vm];
		$ctrl.$onEnter = function (param, rs, rj) {
			//获取分类
			var type = 1;
			switch (mmState.query.type) {
			case 'rpg':
				type = 2;
				break;
			case 'pk':
				type = 3;
				break;
			case 'chess':
				type = 4;
				break;
			case 'other':
				type = 0;
				break;
			}

			//设置当前类型
			vm.type = type;

			//清空内容
			vm.lists.clear();

			mmState.query.from = mmState.query.from || 0;
			mmState.query.number = mmState.query.number || 20;

			//请求获取分页信息
			post('/api/game/getGameList', {
				type: type,
				from: mmState.query.from,
				number: mmState.query.number
			}, null, '获取信息失败：', function (data) {
				for (var key in data.list) {
					data.list[key].GameImage = data.list[key].GameImage == "" ? "/static/img/app.png" : "http://img.wokugame.com/" + data.list[key].GameImage;
				}
				vm.lists = data.list;

				//生成分页
				vm.pagin = createPagin(mmState.query.from, mmState.query.number, data.count);
			});
		}
		$ctrl.$onRendered = function () {}
	});
});