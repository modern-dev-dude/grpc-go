import Component from '@glimmer/component';
import { action } from "@ember/object"

export default class LoginFormComponent extends Component {
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
}

