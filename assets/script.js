// Simple router
function r(l, f) { if(window.location.pathname == l) { f() } }

// Get element
function e(i) { return document.getElementById(i) }

// AJAX request
function d(m, u, d, c) {
  // new instance
  var newXHR = new XMLHttpRequest();

  // set listener
  newXHR.addEventListener('load', c);

  // Set url
  newXHR.open(m, u);

  // Set body format
  newXHR.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

  // Send data
  newXHR.send(d);
}

function a(t, m) {
  // Get header element
  var h = document.getElementsByTagName('header')[0],
      id = Math.random()

  // Clear all and add flash
  h.innerHTML = '<flash ' + t + ' id="flash-' + id + '">' + m + '</flash>'

  setTimeout(function() {
    e('flash-' + id).remove()
  }, 5000)
}

function remove(el) {
  d('DELETE', '/u/a/del/' + el.getAttribute('data-link') + '?token=' + t, null, function() {
    // Parse response
    var response = JSON.parse(this.responseText)

    // Handle error
    if(response['error']) {
      a('error', response['message'])
      return
    }

    // Create alert
    a('info', 'Deleted')

    // Remove parent
    el.parentNode.remove()

    // Check if links are empty
    if(e('list-links').innerText.trim() == '') {
      e('list-links').innerHTML = '<p>No links found</p>'
    }
  })
}

function removeTriggers() {
  var el_delete = document.getElementsByTagName('links-delete')

  for(var i = 0; i < el_delete.length; i++) {
    el_delete[i].onclick = function() {
      if(confirm('are you sure to delete this?')) {
        remove(this)
      }
    }
  }
}

r('/', function() {
  var el_submit = e('b-submit'),
      el_link   = e('i-link')

  // Listen for click event
  el_submit.onclick = function(e) {
    // Prevent from submit
    e.preventDefault()

    // Post request to api with token
    d('POST', '/a/links', 'link=' + encodeURI(el_link.value), function() {
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
    d('POST', '/u/a/links?token=' + t, 'link=' + encodeURI(el_link.value) + '&slug=' + encodeURI(el_slug.value), function() {
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

      // Clear links element if required
      if(el_links.innerText.trim() == 'No links found') {
        el_links.innerHTML = ''
      }

      // Append link
      el_links.innerHTML += '<row>' +
        '  <a href="/l/' + s + '">/' + s + '</a>' +
        '  <visits>0</visits> visits' +
        '  <links-delete class="links-delete" data-link="' + s + '">Delete</links-delete>' +
        '</row>'

      // Links delte action reassign
      removeTriggers()
    })
  }

  // Links delete action
  removeTriggers()
})
