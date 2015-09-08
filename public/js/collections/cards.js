// Collection of Card Models
var CardsCollection = Backbone.Collection.extend({
  model: Card,
  url: '/api/cards',
});

// create global cards collection
var Cards = new CardsCollection();
