var CardView = Backbone.View.extend({

  tagName:  'card',
  cardTpl: _.template( $('#card-template').html() ),

  // add rating event here later
  events: {
    // 'dblclick label': 'edit',
    // 'keypress .edit': 'updateOnEnter',
    // 'blur .edit':   'close'
    'click .card': 'show'
  },

  // Called when the view is first created
  initialize: function() {
    this.listenTo(this.model, 'change', this.render);
    this.listenTo(this.model, 'filter', this.filter);
  },

  filter: function() {
    this.$el.css('display', this.model.attributes.visible);
  },

  // Re-render the titles of the card item.
  render: function() {
    if (this.model.attributes.visible) {
      tmpObject = this.model.attributes;
      tmpObject.hours = this.formatTime(this.model.attributes.CreatedOn);
      this.$el.html( this.cardTpl( tmpObject ) );
    }
    else
      this.$el.html('');
    return this;
  },

  show: function(env) {
    if (this.model.attributes.existingCard) {
      var id = $(env.currentTarget).data('id');
      location = "/cards/" + id;
    }
  },

  formatTime: function(d) {
    var date = new Date(d);
    var h = date.getUTCHours();
    var amPmString = (h <= 11)  ?" am" : " pm";
    h = h % 12;
    var min = date.getUTCMinutes();
    min = (min < 10) ? "0" + min : min;
    h = (h === 0) ? 12 : h;
    return h + ":" + min + amPmString;
  },

  // edit: function() {
    // executed when card label is double clicked
  // },

  // close: function() {
    // executed when card loses focus
  // },

  // updateOnEnter: function( e ) {
    // executed on each keypress when in card edit mode,
    // but we'll wait for enter to get in action
  // }
});
