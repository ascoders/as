'use strict'

define("yuqingList", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "yuqingList",
		categoryName: '',
		lists: [],
		listsRendered: function () {
			jbox()
		},
		debug: false, // 调试模式，显示所有词性
		debugChange: function () {
			setTimeout(function () {
				if (!vm.debug) {
					return
				}
				jbox()
			}, 100)
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			var category = param.category == 'all' ? '' : param.category

			post('/api/yuqing/' + category, {
				from: mmState.query.from || 0,
				number: mmState.query.number || 20
			}, null, '暂无内容', function (data) {
				if (data == null) {
					avalon.router.navigate('/yuqing')
					return
				}

				// 设置页面信息
				switch (param.category) {
				case 'all':
					vm.categoryName = '总体'
					break
				case 'tech':
					vm.categoryName = '科技'
					break
				case 'finance':
					vm.categoryName = '金融'
					break
				case 'military':
					vm.categoryName = '军事'
					break
				case 'sport':
					vm.categoryName = '体育'
					break
				case 'entertainment':
					vm.categoryName = '娱乐'
					break
				case 'education':
					vm.categoryName = '教育'
					break
				case 'test':
					vm.categoryName = '测试表'
					break
				default:
					vm.categoryName = param.category
				}

				document.title = vm.categoryName + ' 舆情分析 - 我酷游戏'

				// 替换正面负面文字信息
				for (var i = 0, j = data.length; i < j; i++) {
					if (data[i].d == null) {
						continue
					}

					// 操作保存在数组里，循环完毕后再增加添加否定组
					var degreeOperateGroup = []

					// 当前否定词性
					var degreeCoef = 1

					// 否定句开始，结束位置
					var degreeStart, degreeEnd = -1

					// 否定词是否被阻断
					var degreeLimit = true

					// 是否遇到一个评价或情感，组成有效否定词组
					var degreeLegal = true

					// 阻断否定
					var doDegreeLimit = function (index, isBreak) {
						if (degreeStart == undefined || degreeLimit) {
							return
						}

						degreeCoef *= data[i].d[m].g
						degreeEnd = index

						if (isBreak) {
							degreeLegal = false
						}

						degreeOperateGroup.push({
							start: degreeStart,
							end: degreeEnd,
							coef: degreeCoef,
							legal: degreeLegal
						})

						// 否定被阻断，否定词性清空
						degreeLimit = true
						degreeCoef = 1
						degreeLegal = true
					}

					for (var m = 0, n = data[i].d.length; m < n; m++) {
						// c:词语 g:好坏 t:类型
						var _type = ''
						var _class = ''

						switch (data[i].d[m].t) {
						case -1:
							data[i].d[m]._type = '断句'
							_class = 'break'

							doDegreeLimit(m, true)
							break
						case 0:
							data[i].d[m]._type = '形容词'
							_class = 'adjective'

							doDegreeLimit(m)
							break
						case 1:
							data[i].d[m]._type = '名词'
							_class = 'noun'

							// 有情感倾向才阻断
							if (data[i].d[m].g !== 0) {
								doDegreeLimit(m)
							}
							break
						case 2:
							data[i].d[m]._type = '动词'
							_class = 'verb'

							// 有情感倾向才阻断
							if (data[i].d[m].g !== 0) {
								doDegreeLimit(m)
							}
							break
						case 3:
							data[i].d[m]._type = '程度副词'
							_class = 'degree'

							break
						case 4:
							data[i].d[m]._type = '否定词'
							_class = 'negavite'

							if (degreeLimit) {
								degreeLimit = false
								degreeStart = m
								degreeEnd = m
							} else {
								degreeEnd = m
							}
							degreeCoef *= data[i].d[m].g
							break
						case 5:
							data[i].d[m]._type = '介词'
							_class = 'preposition'

							break
						case 6:
							data[i].d[m]._type = '语气助词'
							_class = 'auxiliary'

							break
						}

						// 标注好坏的class
						var goodClass = ''

						if (data[i].d[m].g > 0 && (data[i].d[m].t == 0 || data[i].d[m].t == 1 || data[i].d[m] == 2)) {
							goodClass = ' good'
						} else if (data[i].d[m].g < 0 && (data[i].d[m].t == 0 || data[i].d[m].t == 1 || data[i].d[m] == 2)) {
							goodClass = ' bad'
						}

						// 标注出的分词增加class
						if (data[i].d[m].t != 0 && data[i].d[m].t != 4) {
							data[i].w[data[i].d[m].p] = '<span class="word' + goodClass + '" title="' + data[i].d[m]._type + ':' + data[i].d[m].g + '" ms-class="' + _class + ':debug" ms-class-1="jbox:debug">' + data[i].d[m].c + '</span>'
						} else {
							data[i].w[data[i].d[m].p] = '<span class="word' + goodClass + '" title="' + data[i].d[m]._type + ':' + data[i].d[m].g + '" ms-class="jbox ' + _class + '">' + data[i].d[m].c + '</span>'
						}
					}

					for (var m = 0, n = degreeOperateGroup.length; m < n; m++) {
						// 无效否定词组不会显示
						if (!degreeOperateGroup[m].legal) {
							continue
						}

						var start = data[i].d[degreeOperateGroup[m].start].p
						var end = data[i].d[degreeOperateGroup[m].end].p

						var degreeText = ''

						for (var p = degreeOperateGroup[m].start; p < degreeOperateGroup[m].end; p++) {
							// 跳过非否定词
							if (data[i].d[p].t != 4) {
								continue
							}

							degreeText += data[i].d[p].c + ' + '
						}

						degreeText += data[i].d[degreeOperateGroup[m].end].c + ' = ' + data[i].d[degreeOperateGroup[m].end]._type + ':' + degreeOperateGroup[m].coef

						var _class = ''

						if (degreeOperateGroup[m].coef > 0) {
							_class = ' degreeGroupGood'
						} else if (degreeOperateGroup[m].coef < 0) {
							_class = ' degreeGroupBad'
						}

						data[i].w[start] = '<span class="degreeGroup jbox' + _class + '" jbox-position-y="bottom" title="' + degreeText + '">' + data[i].w[start]
						data[i].w[end] = data[i].w[end] + '</span>'
					}

					data[i].word = ''

					for (var p = 0, q = data[i].w.length; p < q; p++) {
						data[i].word += data[i].w[p]
					}
				}

				vm.lists = data || []
			})
		}

		$ctrl.$onRendered = function () {

		}
	});
});