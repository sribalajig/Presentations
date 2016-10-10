var app = app || {};

(function ($) {
	'use strict';

	var proxiedSync = Backbone.sync;
    
	app.AppView = Backbone.View.extend({
		el: '.todoapp',

		initialize: function () {
			app.presentations.fetch({reset: true});
		},

		render: function () {
		}
	});
})(jQuery);