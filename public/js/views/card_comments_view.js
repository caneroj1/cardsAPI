var CardCommentsView = Backbone.View.extend({
  tagName: 'comments',
  model: null,
  commentsTpl: _.template( $('#card-comments-template').html()),

  initialize: function(cardComments) {

  },

  render: function() {
    this.$el.html( this.commentsTpl( this.model.attributes ) );
    return this;
  }
});
