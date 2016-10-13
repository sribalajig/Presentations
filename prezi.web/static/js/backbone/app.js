var app = app || {};

$(function() {
    'use strict';   

    var presentations = new app.Presentations();

    $("#search-container").append(new app.SearchView({ collection: presentations }).render().$el);
    $("#grid").append(new app.Grid(presentations).render().$el);
    $("#paginator").append(new app.Paginate(presentations).render().$el);

    presentations.fetch({
            reset: true
    });
});
