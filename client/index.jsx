import React, { Component, Fragment } from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';

import "bootstrap-css-only/css/bootstrap.css";

class App extends Component {
	render() {
		console.debug('rendering');
		return (
			<div className="container-fluid">
				<h1>Welcome to Brian's Go/React App</h1>
				<form action="login" method="post" className="form-group">
					<div className="form-group">
						<label>Email</label>
						<input type="input" name="email" className="form-control"/>
						<label>Password</label>
						<input type="password" name="password" className="form-control"/>
					</div>
					<input type="submit" value="Login" className="btn btn-primary"/>
				</form>
			</div>
		)
	}
}

ReactDOM.render(<App/>, document.getElementById('app'));

