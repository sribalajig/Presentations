var app = app || {};

(function() {
    app.SearchView = Backbone.View.extend({
        events: {
            "keyup #search": "search"
        },
        render: function(data) {
            $(this.el).html(this.template);
            return this;
        },
        initialize: function() {
            this.template = _.template($("#search-tmpl").html());
            this.collection.bind("reset", this.render, this);
        },
        search: function(e) {
            var letters = $("#search").val();
            this.collection.search(letters);
        }
    });
})()
