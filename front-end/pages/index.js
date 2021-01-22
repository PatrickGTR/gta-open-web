import Header from '../components/header'
import Navbar from '../components/navbar'
import SideBar from '../components/sidebar'

export default function Home() {
  return (
    <>
    <Navbar />
    <div class="container">
        <Header title="GTA Open | Home"></Header>
        <div class="row">
            <div class="column column-60">
                <div class="home-content">
                    <h1>Insert title here</h1>
                    <div class="author-date">
                        <div class="row">
                            <div class="column">
                            Patrick Subang
                            </div>
                            <div class="column column-offset-40">
                            January, 1, 2021 at 4:25 AM
                            </div>
                        </div>
                    </div>
                    <center>

                    <img width="256"src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b3/Responsive_Web_Design_Demo_Template.svg/1024px-Responsive_Web_Design_Demo_Template.svg.png"/>
                    </center>
                    <p>
                        What is Lorem Ipsum?
                        Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                        Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                        when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                        It has survived not only five centuries, but also the leap into electronic typesetting,
                        remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                        sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like
                        Aldus PageMaker including versions of Lorem Ipsum.
                    </p>
                    <p>
                        What is Lorem Ipsum?
                        Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                        Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                        when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                        It has survived not only five centuries, but also the leap into electronic typesetting,
                        remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                        sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like
                        Aldus PageMaker including versions of Lorem Ipsum.
                    </p>
                </div>
            </div>
            <div class="column column-offset-10"><SideBar/></div>
        </div>


    </div>
    <footer class="footer">
        Made with 💖 by Patrick Subang | Copyright 2021 GTA Open
    </footer>
    </>
  )
}
