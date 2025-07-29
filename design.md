### Overview 
Mono repos for front end development over large scale applications is damaging to release schedules. Over the past decade industry has been moving to a "micro front end" architecture. This presents an interesting challenge as we face the complexity of a high distributed front-end world.  

#### Industry standard
Today it feels as if most of the industry has adopted module federation as the defacto micro front end solution. Module federation creates a solution for MFEs to work cross teams with ease. However, there is a technical challenge with MF, it lacks the ability to allow teams to own their own build systems. For teams to own their own build system it would require a great deal of orchestration with remote-entry files and cross team collaboration. 

### Possible alternative
Browsers have adopted [importmaps](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script/type/importmap) over the past few years and it allows for engineers to define common import routes among the page. With this solution we leverage that technology to isolate the location of shared imports like react, vue, or any others needed in client side rendering, hydration, bundle deduplication ...etc.  



### Site level
- CSPs
- Page Shell
  - Header and footer 
- Routing

### individual components 
Each team is responsible for a set of widgets ? 
If so how does that work? 

Concept of building a page
- register a page template
    ```html
    <!--  Owned by y team-->
    <HTML-Shell>
    <!-- Owned by x team-->
    <header></header>
    <!-- Owned n team-->
        <page-widget></page-widget>
    <!-- Owned by z team-->
    <footer></footer>
    </HTML-Shell>
    ```