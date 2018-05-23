// Simple router
function r(l, f) { if(window.location.pathname == l) { f() } }

// Get element
function e(i) { return document.getElementById(i) }

// AJAX request
function d(u, d, c) {
  // new instance
  var newXHR = new XMLHttpRequest();

  // set listener
  newXHR.addEventListener('load', c);

  // Set url
  newXHR.open('POST', u);

  // Set body format
  newXHR.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

  // Send data
  newXHR.send(d);
}

function a(t, m) {
  // Get header element
  var h = document.getElementsByTagName('header')[0]

  // Clear all and add flash
  h.innerHTML = '<flash ' + t + '>' + m + '</flash>'
}

r('/', function() {
  var el_submit = e('b-submit'),
      el_link   = e('i-link')

  // Listen for click event
  el_submit.onclick = function(e) {
    // Prevent from submit
    e.preventDefault()

    // Post request to api with token
    d('/a/links', 'link=' + encodeURI(el_link.value), function() {
      // Parse response
      var response = JSON.parse(this.responseText)

      // Handle error
      if(response['error']) {
        a('error', response['message'])
        return
      }

      // Get slug
      var s = response['data']

      // Create alert
      a('info', '<a href="/l/' + s + '">https://' + window.location.hostname + '/l/' + s + '</a>')

      // Clear input field
      el_link.value = ''
    })
  }
})

// Check user board route
r('/u/board', function() {
  var el_submit = e('b-submit'),
      el_link   = e('i-link'),
      el_slug   = e('i-slug'),
      el_links  = e('list-links')

  // Listen for click event
  el_submit.onclick = function(e) {
    // Prevent form submit
    e.preventDefault()

    // Post request to api with token
    d('/u/a/links?token=' + t, 'link=' + encodeURI(el_link.value) + '&slug=' + encodeURI(el_slug.value), function() {
      // Parse response
      var response = JSON.parse(this.responseText)

      // Handle error
      if(response['error']) {
        a('error', response['message'])
        return
      }

      // Get slug
      var s = response['data']

      // Create alert
      a('info', '<a href="/l/' + s + '">https://' + window.location.hostname + '/l/' +  s + '</a>')

      // Clear input field
      el_link.value = ''
      el_slug.value = ''
    })
  }
})
