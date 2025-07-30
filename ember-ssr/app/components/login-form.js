import Component from '@glimmer/component';
import { action } from "@ember/object"
import { tracked } from "@glimmer/tracking"

export default class LoginFormComponent extends Component {
    @tracked
    userId = null;

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
        
        this.loginAsUserWithId(user)
    }

    loginAsUserWithId(val){
        console.log(val)    
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

