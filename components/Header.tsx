import Image from 'next/image'
import Logo from "@/public/images/logo.png"

export default function Header (){
  return (
    <div className='flex'>
        <div className='flex justify-center'>
            <div>
                <Image alt='logo' src={Logo} />
            </div>
            <p className='text-white'>ShopSphere</p>
        </div>
        <div>

        </div>
    </div>
  )
}
