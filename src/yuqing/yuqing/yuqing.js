'use strict';

define("yuqing", ['jquery', 'jquery.timeago'], function ($) {
	var vm = avalon.define({
		$id: "yuqing",
		$setCharts: function (itemName, datas) {
			require(['echarts'], function (echarts) {
				var myChart = echarts.init(document.getElementById(itemName));

				var items = [{
					key: 'good',
					value: '正面'
				}, {
					key: 'normal',
					value: '中立'
				}, {
					key: 'bad',
					value: '负面'
				}];

				// 转换数据接口
				var xData = [];
				for (var i = 0; i < datas.length; i++) {
					xData.push(datas[i].time);
				}

				var legendData = [];
				for (var i = 0, j = items.length; i < j; i++) {
					legendData.push(items[i].value);
				}

				var series = [];
				for (var i = 0; i < items.length; i++) {
					var obj = {
						name: items[i].value,
						type: 'line',
						smooth: true,
						itemStyle: {
							normal: {
								areaStyle: {
									type: 'default'
								}
							}
						},
						data: []
					};

					for (var j = 0; j < datas.length; j++) {
						obj.data.push(datas[j][items[i].key]);
					}

					series.push(obj);
				}

				// 为echarts对象加载数据 
				myChart.setOption({
					tooltip: {
						trigger: 'axis'
					},
					legend: {
						data: legendData
					},
					toolbox: {
						show: false
					},
					calculable: true,
					xAxis: [{
						type: 'category',
						boundaryGap: false,
						data: xData
					}],
					color: ['#47C947', '#8F8F8F', '#B95959'],
					yAxis: [{
						type: 'value'
					}],
					series: series
				})
			})
		}
	})

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {
			// 获取各表信息
			post('/api/yuqing/charts', null, null, null, function (data) {
				for (var key in data) {
					if (data[key] == null || data[key].length == 0) {
						return
					}

					var datas = []

					for (var i = 0, j = data[key].length; i < j; i++) {
						var obj = {
							good: data[key][i].g,
							normal: data[key][i].n,
							bad: data[key][i].b,
							time: $.timeago(data[key][i].cr)
						}

						datas.push(obj)
					}

					vm.$setCharts(key, datas)
				}
			})

			// 获取中国当日舆情分析
			var chinaGood = []
			var chinaBad = []
			var goodMax = 0
			var badMax = 0
			post('/api/yuqing/china', null, null, null, function (data) {
				for (var key in data) {
					if (data[key] == null) {
						return
					}

					if (data[key].g > 0) {
						chinaGood.push({
							name: key,
							value: data[key].g
						})
					}

					if (data[key].b > 0) {
						chinaBad.push({
							name: key,
							value: -data[key].b
						})
					}

					if (data[key].g > goodMax) {
						goodMax = data[key].g
					}

					if (data[key].b > badMax) {
						badMax = data[key].g
					}
				}

				//设置中国地图实时新闻
				require(['echarts'], function (echarts) {
					var option = {
						title: {
							text: '全国实况',
							subtext: '今日舆情一览',
							x: 'center'
						},
						tooltip: {
							trigger: 'item'
						},
						legend: {
							orient: 'vertical',
							x: 'left',
							data: ['正面', '负面']
						},
						color: ['#47C947', '#B95959'],
						dataRange: {
							min: -badMax,
							max: goodMax,
							x: 'left',
							y: 'bottom',
							text: ['正面', '负面'], // 文本，默认为数值文本
							calculable: true,
							color: ['#47C947', '#8F8F8F', '#B95959']
						},
						toolbox: {
							show: false
						},
						roamController: {
							show: true,
							x: 'right',
							mapTypeControl: {
								'china': true
							}
						},
						series: [{
							name: '正面',
							type: 'map',
							mapType: 'china',
							roam: false,
							itemStyle: {
								normal: {
									label: {
										show: true
									}
								},
								emphasis: {
									label: {
										show: true
									}
								}
							},
							data: chinaGood
						}, {
							name: '负面',
							type: 'map',
							mapType: 'china',
							itemStyle: {
								normal: {
									label: {
										show: true
									}
								},
								emphasis: {
									label: {
										show: true
									}
								}
							},
							data: chinaBad
						}]
					}

					var myChart = echarts.init(document.getElementById('china'));

					myChart.on('click', function (data) {
						avalon.router.navigate('/yuqing/' + data.name)
					})

					myChart.setOption(option)
				})
			})
		}
	});
});