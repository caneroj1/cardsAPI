var NewCardPageView = Backbone.View.extend({
  el: '#new-card',
  cardView: new CardView(),
  color: 0,

  events: {
    'submit #new-card-form':  'submitCard',
    'click #black':           'changeCard',
    'click #white':           'changeCard',
    'keyup #cardBody':        'updateBody'
  },

  initialize: function() {
    this.listenTo(Cards, 'add', this.submitted);
    cardView = new CardView({ model: new Card({existingCard: false}) });
    this.$('#new-card-body').html(cardView.render().el);
  },

  submitted: function(card) {
    console.log("Created card");
    console.log(card);
  },

  submitCard: function() {
    var blanks = this.$('#cardBody').val().split(' ').reduce(function(prev, curr) {
      return /_+[.?!,]?/.test(curr) ? prev + 1 : prev;
    }, 0);

    var newCard = new Card({
      CardType: this.color,
      CardBody: this.$('#cardBody').val(),
      CardBlanks: blanks,
      noMapBody: true
    });

    Cards.create(newCard, {emulateJSON: true});
    return false;
  },

  updateBody: function(env) {
    if (this.color === 0)
      this.$('.cah-face').html(this.$('#cardBody').val());
    else
      this.$('.cah-face').html($.map(this.$('#cardBody').val().split(' '),
      function(val, i) {
        if (/_+[.?!,]?/.test(val))
          return "______" + val;
        else
          return val;
      }).join(' '));
  },

  changeCard: function(env) {
    var color = $(env.currentTarget).data('color');
    this.color = color;
    var body = $.map(this.$('#cardBody').val().split(' '),
    function(val, i) {
      if (color === 1 && /_+[.?!,]?/.test(val))
        return "______" + val;
      else
        return val;
    }).join(' ');

    cardView = new CardView({ model: new Card({existingCard: false, CardType: color, CardBody: body}) });
    this.$('#new-card-body').html(cardView.render().el);
  },
});

var mainView = new NewCardPageView();
