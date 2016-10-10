var app = app || {};

(function () {
	'use strict';

	var Presentations = Backbone.Collection.extend({
		model: app.Presentation,
		url: 'http://localhost:8092/v1/presentations'
	});

	app.presentations = new Presentations();
})();