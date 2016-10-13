var app = app || {};

app.Paginate = function(presentations) {
    return new Backgrid.Extension.Paginator({
        collection: presentations
    });
}
