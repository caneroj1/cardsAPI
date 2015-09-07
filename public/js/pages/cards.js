var CardsPageView = Backbone.View.extend({
  el: '#cards_container',
  type: -1,
  color: -1,

  initialize: function() {
    this.listenTo(Cards, 'add', this.addOne);
    this.listenTo(Cards, 'reset', this.addAll);
  },

  events: {
    'click #all-cards':     'filterByClassic',
    'click #classic-cards': 'filterByClassic',
    'click #user-cards':    'filterByClassic',
    'click #new-cards':     'filterByClassic',
    'click #both-cards':    'filterByColor',
    'click #black-cards':   'filterByColor',
    'click #white-cards':   'filterByColor',
  },

  addOne: function(card) {
    var cardView = new CardView({ model: card });
    this.$('#cards').append( cardView.render().el );
  },

  addAll: function() {
    this.$('#cards').html('');
    Cards.each(this.addOne, this);
  },

  filterByClassic: function(env) {
    this.type = $(env.currentTarget).data('classic');
    var type = this.type;
    var color = this.color;
    Cards.models.forEach(function(card){
      card.filter(type, color);
    })
  },

  filterByColor: function(env) {
    this.color = $(env.currentTarget).data('color');
    var type = this.type;
    var color = this.color;
    Cards.models.forEach(function(card){
      card.filter(type, color);
    })
  },
});

var mainView = new CardsPageView();
