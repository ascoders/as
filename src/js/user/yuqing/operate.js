'use strict';

define("userYuqingOperate", ['jquery', 'js/public/avalon.table'], function ($) {
	var vm = avalon.define({
		$id: "userYuqingOperate",
		segoLoading: true,
		loadSego: function () {
			vm.segoLoading = true
			vm.$freshStatus()
			post('/api/yuqing/loadSego', null, '分词词库正在载入..', '', function () {
				// 加载时间过长，返回只是开始加载信号
			})
		},
		anaylseLoading: true,
		loadAnaylse: function () {
			vm.anaylseLoading = true
			vm.$freshStatus()
			post('/api/yuqing/loadAnaylse', null, '分析词库载入成功', '', function () {
				vm.anaylseLoading = false
			})
		},
		rssLoading: true,
		getRss: function () {
			vm.rssLoading = true
			vm.$freshStatus()
			post('/api/yuqing/getRss', null, '抓取信息成功', '', function () {
				vm.rssLoading = false
			})
		},
		resultLoading: true,
		freshResult: function () {
			vm.resultLoading = true
			vm.$freshStatus()
			post('/api/yuqing/freshResult', null, '结果已更新', '', function () {
				vm.resultLoading = false
			})
		},
		$freshInterval: null,
		$freshStatus: function () {
			if (vm.$freshInterval != null) {
				clearInterval(vm.$freshInterval)
			}

			vm.$freshInterval = setInterval(function () {
				post('/api/yuqing/operateStatus', null, null, '获取操作状态失败：', function (data) {
					vm.segoLoading = data.sego
					vm.anaylseLoading = data.anaylse
					vm.rssLoading = data.rss
					vm.resultLoading = data.result

					if (!vm.segoLoading && !vm.anaylseLoading && !vm.rssLoading && !vm.resultLoading) {
						if (vm.$freshInterval != null) {
							clearInterval(vm.$freshInterval)
						}
					}
				})
			}, 2000)
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			vm.segoLoading = true
			vm.anaylseLoading = true
			vm.rssLoading = true
			vm.resultLoading = true

			post('/api/yuqing/operateStatus', null, null, '获取操作状态失败：', function (data) {
				vm.segoLoading = data.sego
				vm.anaylseLoading = data.anaylse
				vm.rssLoading = data.rss
				vm.resultLoading = data.result

				if (vm.segoLoading || vm.anaylseLoading || vm.rssLoading || vm.resultLoading) {
					vm.$freshStatus()
				}
			})
		}

		$ctrl.$onRendered = function () {

		}
	});
});