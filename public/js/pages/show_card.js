var ShowCardPageView = Backbone.View.extend({
  el: '#show-card',
  tagName: 'card',
  model: new Card({}),
  cardTpl: _.template( $('#show-card-template').html() ),

  initialize: function(card) {
    this.model = card;
    this.model.set({ existingCard: false })
    this.$('#card').html(new CardView( { model: this.model }).render().el );
    this.$('#comments').html(new CardCommentsView( { model: {} }).render().el );
  },

  render: function() {
    this.$('#info').html( this.cardTpl( this.model.attributes ) );
    return this;
  }
});
