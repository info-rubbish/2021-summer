import MarkdownIt from 'markdown-it'
var renderpar = new MarkdownIt({
    linkify: true,
    typographer: true,
})
export default {
    MutiLine(x) {
        // return x
        //     .split(/\r?\n|\r/g)
        //     .map((x) => new MarkdownIt().render(x))
        //     .join('\n')
        return renderpar.render(x)
    },
    SingleLine(x) {
        return renderpar.render(x)
    },
}
