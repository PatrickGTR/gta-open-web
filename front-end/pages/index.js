import SideBar from "../components/sidebar";
import Layout from "../components/layout";

const Home = () => {
  return (
    <Layout title="GTA Open | Home">
      <div className="row">
        <div className="home-content">
          <h1>Insert title here</h1>
          <div className="author-date">
            <div className="row">
              <div className="column">Patrick Subang</div>
              <div className="column column-offset-40">
                January, 1, 2021 at 4:25 AM
              </div>
            </div>
          </div>
          <center>
            <img
              width="256"
              src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b3/Responsive_Web_Design_Demo_Template.svg/1024px-Responsive_Web_Design_Demo_Template.svg.png"
            />
          </center>
          <p>
            What is Lorem Ipsum? Lorem Ipsum is simply dummy text of the
            printing and typesetting industry. Lorem Ipsum has been the
            industry's standard dummy text ever since the 1500s, when an unknown
            printer took a galley of type and scrambled it to make a type
            specimen book. It has survived not only five centuries, but also the
            leap into electronic typesetting, remaining essentially unchanged.
            It was popularised in the 1960s with the release of Letraset sheets
            containing Lorem Ipsum passages, and more recently with desktop
            publishing software like Aldus PageMaker including versions of Lorem
            Ipsum.
          </p>
          <p>
            What is Lorem Ipsum? Lorem Ipsum is simply dummy text of the
            printing and typesetting industry. Lorem Ipsum has been the
            industry's standard dummy text ever since the 1500s, when an unknown
            printer took a galley of type and scrambled it to make a type
            specimen book. It has survived not only five centuries, but also the
            leap into electronic typesetting, remaining essentially unchanged.
            It was popularised in the 1960s with the release of Letraset sheets
            containing Lorem Ipsum passages, and more recently with desktop
            publishing software like Aldus PageMaker including versions of Lorem
            Ipsum.
          </p>
        </div>

        <div className="column column-offset-5">
          <SideBar />
        </div>
      </div>
    </Layout>
  );
};

export default Home;
