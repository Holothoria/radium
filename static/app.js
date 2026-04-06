import { h, render } from 'https://esm.sh/preact@10.19.2'
import { useState}   from 'https://esm.sh/preact@10.19.2/hooks'

const TMDB_KEY = "your_key_here"

function App() {
	return h('div', null,
	  h('nav', {class: 'navbar'},
	   h ('div', {class: 'logo-wrap'},
	     h('div', {class: 'logo'},
			 h('div', {class: 'spoke s1'}),
			 h('div', {class: 'spoke s2'}),
			 h('div', {class: 'spoke s3'}),
			 h('span', {class: 'logo-r'}, 'R')
			),
			
			h('span', {class: 'brand'}, 'RADIUM')
		)
	),
	h('main', null,
	
	  h('div', {class: 'hero'},
	    h('h1', null, 'Watch Everything'),
	    h('p',  null,  'Movies, TV Shows, Anime & More')
	    )
	   )
	  )
	 }
	 
	 render(h(App, null), document.getElementById('app')) 
		  
