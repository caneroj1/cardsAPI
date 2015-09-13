var NewCardPageView = Backbone.View.extend({
  el: '#new-card',
  cardView: new CardView(),
  color: 0,
  blanks: 0,
  failColor: "#ef5350",

  events: {
    'submit #new-card-form':  'submitCard',
    'click #black':           'changeCard',
    'click #white':           'changeCard',
    'keyup #cardBody':        'updateBody'
  },

  initialize: function() {
    this.$('#black-info').css('display', 'none');
    this.listenTo(Cards, 'add', this.submitted);
    cardView = new CardView({ model: new Card({existingCard: false}) });
    this.$('#new-card-body').html(cardView.render().el);
  },

  submitted: function(card) {
    if (!card.isValid())
      this.handleErrors(card.validationError)
    else {
      console.log("SUCCESS SUBMIT");
    }
  },

  handleErrors: function(err) {
    this.$('#fails').html(err).css('color', this.failColor);
  },

  submitCard: function() {
    var blanks = this.$('#cardBody').val().split(' ').reduce(function(prev, curr) {
      return /_+[.?!,]?/.test(curr) ? prev + 1 : prev;
    }, 0);

    var newCard = new Card({
      CardType: this.color,
      CardBody: this.$('#cardBody').val(),
      CardBlanks: blanks,
      CreatorID: 1, // TODO: USE THE USER ID OF THE SIGNED IN USER
      noMapBody: true
    });

    if (newCard.isValid())
      Cards.create(newCard, {emulateJSON: true});
    else
      this.handleErrors(newCard.validationError);
    return false;
  },

  updateBody: function(env) {
    if (this.color === 0) {
      var body = this.$('#cardBody').val();
      if (body.indexOf("_") >= 0)
        this.$('#blanks-counter').html("No underscores allowed.").css("color", this.failColor);
      else
        this.$('#blanks-counter').html("").css("color", "black");
      this.$('.cah-face').html(body);
    }
    else {
      var tBlanks = 0;
      var tFailColor = this.failColor;
      var body = $.map(this.$('#cardBody').val().split(' '),
      function(val, i) {
        if (/_+[.?!,]?/.test(val)) {
          tBlanks += 1;
          if (tBlanks > 3) this.$('#blanks-counter').css("color", tFailColor);
          else this.$('#blanks-counter').css("color", "black");
          this.$('#blanks-counter').html("Blanks: " + tBlanks + "/3");
          return "______" + val;
        }
        else
          return val;
      }).join(' ');
      if (tBlanks == 0)
        this.$('#blanks-counter').css("color", "black");
      this.$('#blanks-counter').html("Blanks: " + tBlanks + "/3");
      this.$('.cah-face').html(body);
      cardView = new CardView({ model: new Card({existingCard: false, CardType: this.color, CardBody: body, CardBlanks: tBlanks}) });
      this.$('#new-card-body').html(cardView.render().el);
    }
  },

  changeCard: function(env) {
    this.$('#fails').html('');
    var color = $(env.currentTarget).data('color');
    this.color = color;
    var tBlanks = 0;
    var tFailColor = this.failColor;
    var body = $.map(this.$('#cardBody').val().split(' '),
    function(val, i) {
      if (color === 1 && /_+[.?!,]?/.test(val)) {
        tBlanks++;
        if (tBlanks > 3) this.$('#blanks-counter').css("color", tFailColor);
        else this.$('#blanks-counter').css("color", "black");
        this.$('#blanks-counter').html("Blanks: " + tBlanks + "/3");
        return "______" + val;
      }
      else {
        return val;
      }
    }).join(' ');

    if (color === 0) {
      if (body.indexOf("_") >= 0)
        this.$('#blanks-counter').html("No underscores allowed.").css("color", this.failColor);
      else
        this.$('#blanks-counter').html("").css("color", "black");
    }
    else {
      if (tBlanks > 3) this.$('#blanks-counter').css("color", this.failColor);
      else this.$('#blanks-counter').css("color", "black");
    }

    if (color == 1) {
      this.$('#white-info').css('display', 'none');
      this.$('#black-info').css('display', 'block');
    }
    else {
      this.$('#white-info').css('display', 'block');
      this.$('#black-info').css('display', 'none');
    }

    cardView = new CardView({ model: new Card({existingCard: false, CardType: this.color, CardBody: body, CardBlanks: tBlanks}) });
    this.$('#new-card-body').html(cardView.render().el);
  },
});

var mainView = new NewCardPageView();
