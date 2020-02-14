import { h } from 'preact'
import { Link } from 'preact-router/match'
import style from './style.css'

// Icons
import iconSyncAlt from '../../assets/fontawesome/sync-alt-solid.svg'
import iconPlus from '../../assets/fontawesome/plus-solid.svg'
import iconCog from '../../assets/fontawesome/cog-solid.svg'

const Footer = () => (
  <div class={style.footer}>
    <img class={style.footerIcon} src={iconSyncAlt} />
    <div className={style.iconContainerMiddle}>
      <img class={style.footerIcon} src={iconPlus} />
    </div>
    <img class={style.footerIcon} src={iconCog} />
    {/* <nav>
      <Link activeClassName={style.active} href="/">Home</Link>
      <Link activeClassName={style.active} href="/profile">Me</Link>
      <Link activeClassName={style.active} href="/profile/john">John</Link>
    </nav> */}
  </div>
)

export default Footer
