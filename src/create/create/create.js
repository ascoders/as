'use strict';
define("create", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "create",
		name: '',
		path: '',
		type: 1,
		submit: function () { //提交
			var teshu = /[`~!！@#$%^&*()_+<>?:"”{},.，。\/;；‘'[\]]/im;
			if (vm.name == "" || teshu.test(vm.name)) { //名称不能为空，不能包含特殊符号
				notice("名称不能或者含有特殊符号", 'red');
				return;
			}
			if (vm.name.length < 2 || vm.name.length > 20) {
				notice("名称长度2-20", 'red');
				return;
			}
			var xiaoxie = /^[a-z]*$/g;
			if (vm.path == "" || !xiaoxie.test(vm.path)) { //域名不能为空，必须为字母
				notice("域名只包括字母", 'red');
				return;
			}
			if (vm.path.length < 3 || vm.path.length > 20) {
				notice("域名长度3-20", 'red');
				return;
			}

			post('/api/game/createGame', {
				name: vm.name,
				path: vm.path,
				type: vm.type,
			}, '创建成功', '', function () {
				//跳转到游戏首页
				avalon.router.navigate('/g/' + vm.path);
			});
		},
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$vmodels = [vm];
		$ctrl.$onEnter = function (param, rs, rj) {
			document.title = '创建网站 - 我酷游戏';
		}
		$ctrl.$onRendered = function () {}
	});
});