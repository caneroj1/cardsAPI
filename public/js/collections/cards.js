// Collection of Card Models
var CardsCollection = Backbone.Collection.extend({
  model: Card,
  url: 'http://localhost:9000/api/cards',
});

// create global cards collection
var Cards = new CardsCollection();
