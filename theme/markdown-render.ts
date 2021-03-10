import { mode } from "@chakra-ui/theme-tools"
import userCustomTheme from "./user-custom"

export default function markdownRender(props) {
    console.log(props)
    return  {
        '.markdown-render': {
           wordBreak: 'break-word',
           '.hljs' : {
               padding: '1rem',
               borderRadius: '8px'
           },
           'ul,ol' : {
               paddingLeft: '1rem',
               margin: '1.2rem 0',
               li: {
                   margin: '.8rem 0'
               }
           },
           'h1': {
               fontSize: '2rem',
               fontWeight: 'bold',
               marginBottom: '0.8rem'
           },
           'h2': {
            fontSize: '1.8rem',
            fontWeight: 'bold',
            marginBottom: '0.6rem'
            },
            'h3': {
                fontSize: '1.6em',
                fontWeight: '600',
                marginBottom: '0.4rem'
            },
            'h4': {
                fontSize: '1.4em',
                fontWeight: '600'
            },
            'h5,h6': {
                fontSize: '1.2em',
                fontWeight: 'normal'
            },
            p: {
                margin: '1.2rem 0',
            },
            blockquote: {
                lineHeight: '2rem',
                margin: '1.5rem 0',
                p :{
                    paddingLeft: '1rem',
                    fontWeight: '500',
                    fontStyle: 'italic',
                    borderLeftWidth: '.25rem',
                    borderLeftColor: '#e5e7eb',
                    color: mode("inherit", "'rgb(189, 189, 189)'")(props),
                    fontSize: '1.2rem',
                }
            },
            pre: {
                margin: '1.6rem 0',
                fontSize: '.95rem'
            },
            a: {
                textDecoration: 'underline !important'
            },
            '.at-user-link': {
                textDecoration: 'none !important',
                borderBottom: '2px dashed rgb(158, 158, 158)',
                margin:'0 3px',
                color: mode(props.theme.colors.teal[600], props.theme.colors.teal[200])(props),
                fontWeight: '550',
                paddingBottom: '2px'
            }       
        }
    }
}