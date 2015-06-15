'use strict';

define("notFound", ['jquery'], function ($) {
	return avalon.define({
		$id: "notFound",
		onChange: function (state, done) {

		},
		onAfterLoad: function () {
			document.title = '404';
		},
		$skipArray: ['onChange', 'onAfterLoad'],
	});
});