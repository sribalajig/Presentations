var app = app || {};

(function ($) {
	'use strict';

	app.PresentationView = Backbone.View.extend({
		tagName:  'li',

		template: _.template($('#item-template').html()),

		render: function () {
			this.$el.html(this.template(this.model.toJSON()));
			return this;
		},
		
		clear: function () {
			this.model.destroy();
		}
	});
})(jQuery);
