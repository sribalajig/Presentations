var app = app || {};

(function () {
	'use strict';

	app.Presentation = Backbone.Model.extend({
		defaults: {
			id: "",
			title: "Awesome presentation",
			thumbNail: "https://placeimg.com/400/400/any",
			createdAt: "Jan 1, 2016"
		}
	});
})();