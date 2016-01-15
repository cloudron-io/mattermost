// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import * as Client from '../../utils/client.jsx';
import * as AsyncClient from '../../utils/async_client.jsx';

export default class OAuthSettings extends React.Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);

        this.state = {
            Enable: this.props.config.OAuthSettings.Enable,
            saveNeeded: false,
            serverError: null
        };
    }

    handleChange(action) {
        var s = {saveNeeded: true, serverError: this.state.serverError};

        if (action === 'EnableTrue') {
            s.Enable = true;
        }

        if (action === 'EnableFalse') {
            s.Enable = false;
        }

        this.setState(s);
    }

    handleSubmit(e) {
        e.preventDefault();
        $('#save-button').button('loading');

        var config = this.props.config;
        config.OAuthSettings.Enable = ReactDOM.findDOMNode(this.refs.Enable).checked;
        config.OAuthSettings.Secret = ReactDOM.findDOMNode(this.refs.Secret).value.trim();
        config.OAuthSettings.Id = ReactDOM.findDOMNode(this.refs.Id).value.trim();
        config.OAuthSettings.AuthEndpoint = ReactDOM.findDOMNode(this.refs.AuthEndpoint).value.trim();
        config.OAuthSettings.TokenEndpoint = ReactDOM.findDOMNode(this.refs.TokenEndpoint).value.trim();
        config.OAuthSettings.UserApiEndpoint = ReactDOM.findDOMNode(this.refs.UserApiEndpoint).value.trim();

        Client.saveConfig(
            config,
            () => {
                AsyncClient.getConfig();
                this.setState({
                    serverError: null,
                    saveNeeded: false
                });
                $('#save-button').button('reset');
            },
            (err) => {
                this.setState({
                    serverError: err.message,
                    saveNeeded: true
                });
                $('#save-button').button('reset');
            }
        );
    }

    render() {
        var serverError = '';
        if (this.state.serverError) {
            serverError = <div className='form-group has-error'><label className='control-label'>{this.state.serverError}</label></div>;
        }

        var saveClass = 'btn';
        if (this.state.saveNeeded) {
            saveClass = 'btn btn-primary';
        }

        return (
            <div className='wrapper--fixed'>

                <h3>{'Qt Settings'}</h3>
                <form
                    className='form-horizontal'
                    role='form'
                >

                    <div className='form-group'>
                        <label
                            className='control-label col-sm-4'
                            htmlFor='Enable'
                        >
                            {'Enable Sign Up With OAuth: '}
                        </label>
                        <div className='col-sm-8'>
                            <label className='radio-inline'>
                                <input
                                    type='radio'
                                    name='Enable'
                                    value='true'
                                    ref='Enable'
                                    defaultChecked={this.props.config.OAuthSettings.Enable}
                                    onChange={this.handleChange.bind(this, 'EnableTrue')}
                                />
                                    {'true'}
                            </label>
                            <label className='radio-inline'>
                                <input
                                    type='radio'
                                    name='Enable'
                                    value='false'
                                    defaultChecked={!this.props.config.OAuthSettings.Enable}
                                    onChange={this.handleChange.bind(this, 'EnableFalse')}
                                />
                                    {'false'}
                            </label>
                            <p className='help-text'>
                                {'When true, Mattermost allows team creation and account signup using OAuth OAuth.'} <br/>
                            </p>
                            <div className='help-text'>
                                <ol>
                                    <li>{'Log in to your OAuth account and go to Applications -> Profile Settings.'}</li>
                                    <li>{'Enter Redirect URIs "<your-mattermost-url>/login/OAuth/complete" (example: http://localhost:8065/login/OAuth/complete) and "<your-mattermost-url>/signup/OAuth/complete". '}</li>
                                    <li>{'Then use "Secret" and "Id" fields from OAuth to complete the options below.'}</li>
                                    <li>{'Complete the Endpoint URLs below. '}</li>
                                </ol>
                            </div>
                        </div>
                    </div>

                    <div className='form-group'>
                        <label
                            className='control-label col-sm-4'
                            htmlFor='Id'
                        >
                            {'Id:'}
                        </label>
                        <div className='col-sm-8'>
                            <input
                                type='text'
                                className='form-control'
                                id='Id'
                                ref='Id'
                                placeholder='Ex "jcuS8PuvcpGhpgHhlcpT1Mx42pnqMxQY"'
                                defaultValue={this.props.config.OAuthSettings.Id}
                                onChange={this.handleChange}
                                disabled={!this.state.Enable}
                            />
                            <p className='help-text'>{'Obtain this value via the instructions above for logging into OAuth'}</p>
                        </div>
                    </div>

                    <div className='form-group'>
                        <label
                            className='control-label col-sm-4'
                            htmlFor='Secret'
                        >
                            {'Secret:'}
                        </label>
                        <div className='col-sm-8'>
                            <input
                                type='text'
                                className='form-control'
                                id='Secret'
                                ref='Secret'
                                placeholder='Ex "jcuS8PuvcpGhpgHhlcpT1Mx42pnqMxQY"'
                                defaultValue={this.props.config.OAuthSettings.Secret}
                                onChange={this.handleChange}
                                disabled={!this.state.Enable}
                            />
                            <p className='help-text'>{'Obtain this value via the instructions above for logging into OAuth.'}</p>
                        </div>
                    </div>

                    <div className='form-group'>
                        <label
                            className='control-label col-sm-4'
                            htmlFor='AuthEndpoint'
                        >
                            {'Auth Endpoint:'}
                        </label>
                        <div className='col-sm-8'>
                            <input
                                type='text'
                                className='form-control'
                                id='AuthEndpoint'
                                ref='AuthEndpoint'
                                placeholder='Ex ""'
                                defaultValue={this.props.config.OAuthSettings.AuthEndpoint}
                                onChange={this.handleChange}
                                disabled={!this.state.Enable}
                            />
                            <p className='help-text'>{'Enter https://<your-OAuth-url>/oauth/authorize (example https://example.com:3000/oauth/authorize).   Make sure you use HTTP or HTTPS in your URL depending on your server configuration.'}</p>
                        </div>
                    </div>

                    <div className='form-group'>
                        <label
                            className='control-label col-sm-4'
                            htmlFor='TokenEndpoint'
                        >
                            {'Token Endpoint:'}
                        </label>
                        <div className='col-sm-8'>
                            <input
                                type='text'
                                className='form-control'
                                id='TokenEndpoint'
                                ref='TokenEndpoint'
                                placeholder='Ex ""'
                                defaultValue={this.props.config.OAuthSettings.TokenEndpoint}
                                onChange={this.handleChange}
                                disabled={!this.state.Enable}
                            />
                            <p className='help-text'>{'Enter https://<your-OAuth-url>/oauth/token.   Make sure you use HTTP or HTTPS in your URL depending on your server configuration.'}</p>
                        </div>
                    </div>

                    <div className='form-group'>
                        <label
                            className='control-label col-sm-4'
                            htmlFor='UserApiEndpoint'
                        >
                            {'User API Endpoint:'}
                        </label>
                        <div className='col-sm-8'>
                            <input
                                type='text'
                                className='form-control'
                                id='UserApiEndpoint'
                                ref='UserApiEndpoint'
                                placeholder='Ex ""'
                                defaultValue={this.props.config.OAuthSettings.UserApiEndpoint}
                                onChange={this.handleChange}
                                disabled={!this.state.Enable}
                            />
                            <p className='help-text'>{'Enter https://<your-OAuth-url>/api/v3/user.   Make sure you use HTTP or HTTPS in your URL depending on your server configuration.'}</p>
                        </div>
                    </div>

                    <div className='form-group'>
                        <div className='col-sm-12'>
                            {serverError}
                            <button
                                disabled={!this.state.saveNeeded}
                                type='submit'
                                className={saveClass}
                                onClick={this.handleSubmit}
                                id='save-button'
                                data-loading-text={'<span class=\'glyphicon glyphicon-refresh glyphicon-refresh-animate\'></span> Saving Config...'}
                            >
                                {'Save'}
                            </button>
                        </div>
                    </div>

                </form>
            </div>
        );
    }
}

//config.OAuthSettings.Scope = ReactDOM.findDOMNode(this.refs.Scope).value.trim();
//  <div className='form-group'>
//     <label
//         className='control-label col-sm-4'
//         htmlFor='Scope'
//     >
//         {'Scope:'}
//     </label>
//     <div className='col-sm-8'>
//         <input
//             type='text'
//             className='form-control'
//             id='Scope'
//             ref='Scope'
//             placeholder='Not currently used by OAuth. Please leave blank'
//             defaultValue={this.props.config.OAuthSettings.Scope}
//             onChange={this.handleChange}
//             disabled={!this.state.Allow}
//         />
//         <p className='help-text'>{'This field is not yet used by OAuth OAuth. Other OAuth providers may use this field to specify the scope of account data from OAuth provider that is sent to Mattermost.'}</p>
//     </div>
// </div>

OAuthSettings.propTypes = {
    config: React.PropTypes.object
};
