import React, { Component } from 'react';
import PropTypes from 'prop-types';

import "./login-form-styles.less";

export default class LoginForm extends Component {
	render() {
		return (
			<form action="login" method="post" className="form-group login-form">
				<div className="form-group">
					<label>Email</label>
					<input type="input" name="email" className="form-control"/>
					<label>Password</label>
					<input type="password" name="password" className="form-control"/>
				</div>
				<input type="submit" value="Login" className="btn btn-primary"/>
			</form>
		);
	}
}
