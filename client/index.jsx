import React, { Component, Fragment } from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';

class App extends Component {
	render() {
		console.debug('rendering');
		return (
			<div>
				<h1>Welcome to Brian's Home Page</h1>
				<form action="login" method="post">
					<label>Email</label>
					<input type="input" name="email" />
					<label>Password</label>
					<input type="password" name="password" />
					<input type="submit" value="login" />
				</form>
			</div>
		)
	}
}

ReactDOM.render(<App/>, document.getElementById('app'));

