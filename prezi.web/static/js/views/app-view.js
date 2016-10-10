var app = app || {};

(function ($) {
	'use strict';

	var proxiedSync = Backbone.sync;
    
	app.AppView = Backbone.View.extend({
		el: '.prezi-app',

		initialize: function () {			
			this.$list = $('.prezi-app');

			this.listenTo(app.presentations, 'reset', this.addAll);

			app.presentations.fetch({reset: true});
		},

		render: function () {
		},

		addOne: function (presentation) {
			var view = new app.PresentationView({ model: presentation });
			this.$list.append(view.render().el);
		},

		addAll: function () {
			this.$list.html('');
			app.presentations.each(this.addOne, this);
		},

	});
})(jQuery);