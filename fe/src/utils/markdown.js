import MarkdownIt from 'markdown-it'
// var render=
export default {
    MutiLine(x){
        return x.split(/\r?\n|\r/g).map((x)=>new MarkdownIt().render(x)).join('\n');
    }
    ,SingleLine(x ){
        return render(x)
    }
}