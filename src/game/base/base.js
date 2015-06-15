'use strict';

define("gameBase", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "gameBase",
		state: '', //状态
		game: {}, //游戏信息
		category: '', //记录当前分类路径
		categorys: [], //分类信息 news整合在其中
		categoryMenu: {}, //分类菜单，可能会分开显示
		baseRouter: '', //当前所在位置
		menuRendered: function () { // 菜单渲染完毕
			console.log($('#game-base #menu-content').height() / 82);
		}
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			post('/api/game/getInfo', {
				game: param.game
			}, null, '该板块不存在', function (data) {
				//游戏信息
				data.Game.Logo = data.Game.GameImage == "" ? "/static/img/app.png" : "http://img.wokugame.com/" + data.Game.GameImage;
				data.Game.GameImage = data.Game.GameImage == "" ? "" : "http://img.wokugame.com/" + data.Game.GameImage;
				data.Game.Icon = data.Game.Icon == "" ? "" : "http://img.wokugame.com/" + data.Game.Icon;

				//应用截图数量固定为6个，不足补全
				data.Game.Image = data.Game.Image || [];
				for (var key in data.Game.Image) {
					data.Game.Image[key] = "http://img.wokugame.com/" + data.Game.Image[key];
				}
				vm.game = data.Game || [];

				//创建icon
				if (data.Game.Icon != "") {
					//删除旧icon
					$("[rel='shortcut icon']").remove();

					var link = $("<link/>");
					link.attr("rel", "shortcut icon")
						.attr("href", data.Game.Icon)
						.appendTo('head');
				}

				// 分类
				for (var key in data.Categorys) {
					// 设置分类url
					switch (data.Categorys[key].Type) {
					case 0:
						data.Categorys[key]._type = '';
						break;
					case 1:
						data.Categorys[key]._type = '/doc';
						break;
					default:
						data.Categorys[key]._type = '';
						break;
					}

					// 每个分类后附加最新文章
					for (var k in data.News[key]) {
						if (data.News[key][k] == null) {
							continue;
						}

						if (data.News[key][k]["Title"] == "") {
							data.News[key][k]["Title"] = subStr(data.News[key][k]["Content"], 0, 25);
						}
					}
					data.Categorys[key].News = data.News[key] || [];
				}

				// category根据优先级排序
				data.Categorys.sort(function (a, b) {
					return a.RecommendPri > b.RecommendPri ? 1 : -1
				});

				vm.categorys = data.Categorys || [];

				// 恢复链
				rs();
			}, function () {
				//回到首页
				avalon.router.navigate('/');
			});

			// 停止链
			return false;
		}
		$ctrl.$onRendered = function () {}
	});
});