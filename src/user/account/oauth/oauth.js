'use strict';

define("userAccountOauth", ['jquery', 'jquery.timeago'], function ($) {
	// frontia不支持ie7及以下浏览器
	if (ieVersion() != false && ieVersion() <= 7) {
		notice('不支持ie7以下浏览器！', 'red')
		return;
	}

	var vm = avalon.define({
		$id: "userAccountOauth",
		lists: [{
			icon: 'baidu',
			name: 'baidu',
			image: '',
			time: '',
			bind: false,
			old: false
		}, {
			icon: 'qzone',
			name: 'qqdenglu',
			image: '',
			time: '',
			bind: false,
			old: false
		}, {
			icon: 'tsina',
			name: 'sinaweibo',
			image: '',
			time: '',
			bind: false,
			old: false
		}, {
			icon: 'renren',
			name: 'renren',
			image: '',
			time: '',
			bind: false,
			old: false
		}, {
			icon: 'tqq',
			name: 'qqweibo',
			image: '',
			time: '',
			bind: false,
			old: false
		}, {
			icon: 'kaixin',
			name: 'kaixin',
			image: '',
			time: '',
			bind: false,
			old: false
		}],
		image: '', //当前使用头像
		submit: function (name) {
			require(['frontia'], function (frontia) {
				// 初始化登录的配置
				var options = {
					response_type: 'token',
					media_type: name,
					redirect_uri: 'http://www.wokugame.com/oauth',
					client_type: 'web'
				};
				// 登录
				frontia.social.login(options);
			})
		},
		changeImage: function () { // 修改头像
			avalon.nextTick(function () {
				post('/api/user/changeImage', {
					image: vm.image
				}, '头像已更新', '', function () {
					avalon.vmodels.global.my.image = vm.image;
				});
			});
		},
		rendered: function () { // 渲染完毕
			//友好时间
			$(".timeago").timeago();
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			require(['frontia'], function (frontia) {
				// API key 从应用信息页面获取
				var AK = 'RqeMWD9G1m8agmxfj6ngCKRG';

				// 初始化 frontia
				frontia.init(AK);

				$.when(avalon.vmodels.global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
					//获取当前绑定状态
					post('/api/user/oauthList', {}, null, '', function (data) {
						//更新绑定状态
						for (var key in vm.lists.$model) {
							for (var _key in data) {
								if (data[_key].Type == vm.lists[key].name) {
									vm.lists[key].time = data[_key].ExpiresTime;
									vm.lists[key].image = data[_key].Image;
									var timeUnix = (Date.parse(data[_key].ExpiresTime) - Date.parse(new Date())); //距离现在时间戳（秒）
									if (timeUnix < 0) {
										vm.lists[key].old = true;
									}
									vm.lists[key].bind = true;
									break;
								}
							}
						}

						//更新头像状态
						for (var key in vm.lists.$model) {
							if (vm.lists[key].image == avalon.vmodels.global.my.image) {
								vm.image = vm.lists[key].image;
								break;
							}
						}
					});
				});
			})
		}

		$ctrl.$onRendered = function () {

		}
	});
});