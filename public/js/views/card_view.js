var CardView = Backbone.View.extend({

  tagName:  'card',
  cardTpl: _.template( $('#card-template').html() ),

  // add rating event here later
  events: {
    // 'dblclick label': 'edit',
    // 'keypress .edit': 'updateOnEnter',
    // 'blur .edit':   'close'
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
    if (this.model.attributes.visible)
      this.$el.html( this.cardTpl( this.model.attributes ) );
    else
      this.$el.html('');
    return this;
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
