// Define the Card model
var Card = Backbone.Model.extend({
  // Default card attribute values
  defaults: {
    CardBody: '',
    CardType: 0,
    CardBlanks: 0,
    Classic: false,
    CreatedOn: null,
    ModifiedOn: null,
    CreatorID: 0,
    Approved: false,
    Raters: 0,
    Rating: 0.0,
    ID: 0,
    visible: 'block',
    existingCard: true,
    noMapBody: false,
  },

  initialize: function(attributes) {
    if (attributes.noMapBody) return;
    else {
      if (attributes.CardType == 1 && attributes.CardBlanks > 0) {
        attributes.CardBody = $.map(attributes.CardBody.split(' '),
        function(val, i) {
          if (/_+[.?!,]?/.test(val))
            return "______" + val;
          else
            return val;
        }).join(' ');
      }
      this.set({CardBody: attributes.CardBody}, {silent: true});
    }
  },

  validate: function(attributes) {
    if (attributes.CreatorID < 1)
      return "The card must have a valid user ID."
    if (attributes.CardBody === undefined || attributes.CardBody === "")
      return "The card body is required.";
    if (attributes.CardType > 1 || attributes.CardType < 0)
      return "The card type can only be 0, for white cards, or 1, for black cards.";
    if (attributes.CardType == 1) {
      if (attributes.CardBlanks < 0 || attributes.CardBlanks > 3) {
        return "The number of blanks must be in the range of 0 - 3.";
      }
    }
    else {
      if (attributes.CardBlanks != 0) {
        return "There cannot be blanks on a white card.";
      }
    }
  },

  filter: function(classic, color) {
    result = false;
    switch(classic) {
      case -1:
        result = true;
        break;
      case 0:
        result = this.get('Classic');
        break;
      case 1:
        result = !this.get('Classic') && this.get('Approved');
        break;
      case 2:
        result = !this.get('Classic') && !this.get('Approved');
        break;
    }

    switch(color) {
      case -1:
        result = result && true;
        break;
      default:
        result = result && this.get('CardType') === color;
        break;
    }

    if (result)
      this.set({visible: 'block'}, {silent: true});
    else
      this.set({visible: 'none'}, {silent: true});
    this.trigger('filter');
  },
});
