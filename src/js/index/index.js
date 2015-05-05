'use strict';

define("index", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "index",
		games: [],
		tops: [],
		news: [],
		newgames: [],
		hots: [],
		name: '', //新增名称
		path: '', //路径
		type: 1, //类型
		submit: function () { //提交
			var teshu = /[`~!！@#$%^&*()_+<>?:"”{},.，。\/;；‘'[\]]/im;
			if (avalon.vmodels.index.name == "" || teshu.test(avalon.vmodels.index.name)) { //名称不能为空，不能包含特殊符号
				notice("名称不能或者含有特殊符号", 'red');
				return;
			}
			if (avalon.vmodels.index.name.length < 2 || avalon.vmodels.index.name.length > 20) {
				notice("名称长度2-20", 'red');
				return;
			}
			var xiaoxie = /^[a-z]*$/g;
			if (avalon.vmodels.index.path == "" || !xiaoxie.test(avalon.vmodels.index.path)) { //域名不能为空，必须为字母
				notice("域名只包括字母", 'red');
				return;
			}
			if (avalon.vmodels.index.path.length < 3 || avalon.vmodels.index.path.length > 20) {
				notice("域名长度3-20", 'red');
				return;
			}

			post('/api/game/createGame', {
				name: avalon.vmodels.index.name,
				path: avalon.vmodels.index.path,
				type: avalon.vmodels.index.type
			}, '创建成功', '', function () {
				//跳转到游戏首页
				avalon.router.navigate('/g/' + avalon.vmodels.index.path);
			});
		}
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$vmodels = [vm];
		$ctrl.$onEnter = function (param, rs, rj) {
			document.title = '我酷游戏';

			//获取信息
			post('/api/index/getContent', null, null, '获取基础数据失败', function (data) {
				//最火游戏
				for (var key in data.Games) {
					if (data.Games[key].GameImage === "") {
						data.Games[key].GameImage = "/static/img/app.png";
					} else {
						data.Games[key].GameImage = "http://img.wokugame.com/" + data.Games[key].GameImage;
					}
				}
				avalon.vmodels.index.games = data.Games || [];

				//最新资讯
				avalon.vmodels.index.tops = data.Tops || [];

				//最新游戏
				for (var key in data.NewGames) {
					if (data.NewGames[key].GameImage === "") {
						data.NewGames[key].GameImage = "/static/img/app.png";
					} else {
						data.NewGames[key].GameImage = "http://img.wokugame.com/" + data.NewGames[key].GameImage;
					}
				}
				avalon.vmodels.index.newgames = data.NewGames || [];

				//本周热帖
				for (var key in data.HotTopics) {
					data.HotTopics[key].AuthorImage = userImage(data.HotTopics[key].AuthorImage);
				}
				avalon.vmodels.index.hots = data.HotTopics || [];
			});
		}
		$ctrl.$onRendered = function () {}
	});
});