import Component from '@glimmer/component';
import { action } from "@ember/object"
import { tracked } from "@glimmer/tracking"
import { inject as service  } from '@ember/service';
import AuthService from 'shlack/services/auth';


export default class LoginFormComponent extends Component {
    @tracked
    userId = null;
    /**
     * @type {AuthService}
     */
    @service auth
    /**
     * @returns {boolean}
     */
    get isDisabled(){
        return !this.userId
    }
    
    /**
     * 
     * @param {Event & { target: HTMLFormElement}} evt 
     */
    @action 
    onLoginFormSubmit(evt) {
        evt.preventDefault();

        const {target} = evt;
        const user = target.querySelector('select').value
        
        this.auth.loginWithUserId(user)
    }

    /**
     * 
     * @param {Event & {target: HTMLSelectElement}} e 
     */
    @action
    selectChanged(e) {
        this.userId = e.target.value
    }
}

