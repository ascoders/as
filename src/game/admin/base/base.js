'use strict';

define("gameAdminBase", ['jquery'], function ($) {

	var vm = avalon.define({
		$id: "gameAdminBase",
		baseSubmit: function () { //基本信息提交
			var _this = $(this);

			// 按钮处于loading状态
			_this.text('请等待..').attr('disabled', 'disabled');

			post('/api/game/baseSave', {
				game: avalon.vmodels.gameBase.game.Id,
				size: avalon.vmodels.gameBase.game.Size,
				version: avalon.vmodels.gameBase.game.Version,
				need: avalon.vmodels.gameBase.game.Need,
				description: avalon.vmodels.gameBase.game.Description,
				download: avalon.vmodels.gameBase.game.Download
			}, '已保存', '', function () {
				// 按钮复原
				_this.text('保存').removeAttr('disabled');
			}, function () {
				// 按钮复原
				_this.text('保存').removeAttr('disabled');
			});
		},
		baseSetDropzone: function () { //基本信息截图上传
			//为每个截图位置绑定上传事件（截图是动态生成，所以100毫秒后才能取到dom元素）
			setTimeout(function () {
				$('.j-screenshot').each(function () {
					//不会重复为对象加上dropzone
					if (this.dropzone) {
						return;
					}

					var _this = $(this);
					createDropzone(this, 'http://upload.qiniu.com', {
						type: 'gameScreenShot'
					}, ".jpg,.jpeg,.png,.gif", function (data) {
						post('/api/game/uploadHandle', {
							type: data.type,
							game: avalon.vmodels.gameBase.game.Id,
							etag: data.etag, // 已上传的文件名
							ext: data.ext, //后缀名
							name: data.name, //文件路径
							position: parseInt(_this.attr('position'))
						}, '数据处理成功', '', function (data) {
							if (parseInt(_this.attr('position')) + 1 > avalon.vmodels.gameBase.game.Image.length) { //追加
								avalon.vmodels.gameBase.game.Image.push("http://img.wokugame.com/" + data);
							} else { //修改
								avalon.vmodels.gameBase.game.Image.set(_this.attr('position'), "http://img.wokugame.com/" + data);
							}
							vm.baseSetDropzone();
						});
					});
				});
			}, 100);
		}
	});

	return avalon.controller(function ($ctrl) {

		$ctrl.$onEnter = function (param, rs, rj) {
			// 改变页面titile
			document.title = '基础配置 - 管理 - ' + avalon.vmodels.gameBase.game.Name;

			// 设置状态
			avalon.vmodels.gameAdmin.info = 'base';
		}

		$ctrl.$onRendered = function () {
			//应用图标 dropzone
			createDropzone($('#gameImage')[0], 'http://upload.qiniu.com', {
				type: 'gameImage'
			}, ".jpg,.jpeg,.png,.gif", function (data) {
				avalon.vmodels.gameBase.game.GameImage = 'http://img.wokugame.com/' + data.name;
				post('/api/game/uploadHandle', {
					type: data.type,
					game: avalon.vmodels.gameBase.game.Id,
					etag: data.etag, // 已上传的文件名
					ext: data.ext, //后缀名
					name: data.name //文件路径
				}, '数据处理成功', '', function (data) {
					avalon.vmodels.gameBase.game.GameImage = "http://img.wokugame.com/" + data;
				});
			});

			//应用icon dropzone
			createDropzone($('#gameIcon')[0], 'http://upload.qiniu.com', {
				type: 'gameIcon'
			}, ".ico", function (data) {
				avalon.vmodels.gameBase.game.Icon = 'http://img.wokugame.com/' + data.name;
				post('/api/game/uploadHandle', {
					type: data.type,
					game: avalon.vmodels.gameBase.game.Id,
					etag: data.etag, // 已上传的文件名
					ext: data.ext, //后缀名
					name: data.name //文件路径
				}, '数据处理成功', '', function (data) {
					avalon.vmodels.gameBase.game.Icon = "http://img.wokugame.com/" + data;
				});
			});

			//截图dropzone
			vm.baseSetDropzone();

			// tooltip插件
			jbox();
		}

	});

});